from django.test import TestCase, Client, RequestFactory
from django.urls import resolve
from snippets.models import Snippet
from snippets.views import top, snippet_new, snippet_edit, snippet_detail
from django.contrib.auth import get_user_model

UserModel = get_user_model()


class TopPageViewTest(TestCase):
    def test_top_page_returns_200_and_expected_title(self):
        response = self.client.get('/')
        self.assertContains(response, 'Django Snippets', status_code=200)

    def test_top_page_uses_expected_template(self):
        response = self.client.get('/')
        self.assertTemplateUsed(response, 'snippets/top.html')


class CreateSnipppetTest(TestCase):
    def test_should_resolve_snippet_new(self):
        found = resolve('/snippets/new/')
        self.assertEqual(snippet_new, found.func)


class SnippetDetailTest(TestCase):
    def setUp(self):
        self.user = UserModel.objects.create(
            username='test_user',
            email='test@test.com',
            password='detail_secret_pass0002',
        )
        self.snippet = Snippet.objects.create(
            title='test title',
            code='print("test")',
            description='test description',
            created_by=self.user,
        )

    def test_should_use_expected_template(self):
        response = self.client.get('/snippets/%s/' % self.snippet.id)
        self.assertTemplateUsed(response, 'snippets/snippet_detail.html')

    def test_top_page_returns_200_and_expected_heading(self):
        response = self.client.get('/snippets/%s/' % self.snippet.id)
        self.assertContains(response, self.snippet.title, status_code=200)


class EditSnippetTest(TestCase):
    def test_should_resolve_snippet_edit(self):
        found = resolve('/snippets/1/edit/')
        self.assertEqual(snippet_edit, found.func)


class TopPageRenderSnippetsTest(TestCase):
    def setUp(self):
        self.user = UserModel.objects.create(
            username='test_user',
            email='test@test.com',
            password='top_secret_pass0001',
        )
        self.snippet = Snippet.objects.create(
            title='test title',
            code='print("test")',
            description='test description',
            created_by=self.user,
        )
    
    def test_should_return_snippet_title(self):
        request = RequestFactory().get('/')
        request.user = self.user
        response = top(request)
        self.assertContains(response, self.snippet.title)

    def test_should_return_username(self):
        request = RequestFactory().get('/')
        request.user = self.user
        response = top(request)
        self.assertContains(response, self.user.username)


class CreateSnippetTest(TestCase):
    def setUp(self):
        self.user = UserModel.objects.create(
            username='test_user',
            email='test@test.com',
            password='top_secret_pass0001',
        )
        self.client.force_login(self.user)

    def test_render_creation_form(self):
        response = self.client.get('/snippets/new/')
        self.assertContains(response, 'Create Snippet', status_code=200)

    def test_create_snippet(self):
        data = {'title':'title', 'code':'code', 'description':'description'}
        self.client.post('/snippets/new/', data)
        snippet = Snippet.objects.get(title='title')
        self.assertEqual('code', snippet.code)
        self.assertEqual('description', snippet.description)
