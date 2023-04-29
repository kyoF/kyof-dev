// const printName = (myName: string) => {
// 	console.log(`Hello ${myName}!`);
// };

// const myName: string = 'kyosuke';

// printName(myName);
// printName(1);


// const message = 'Hello!';
// message();

// for (let i=0; i<10; i++) {
//  	console.log(i);
// };

// console.log(i);

// const array: string[] = [];
// const array2: string[] = [...array, 'string'];
// console.log(array2);

// const testFunc = (obj: { firstName: string, lastName?: string }) => {
// 	return `${obj.firstName} ${obj.lastName}`;
// }
// console.log(testFunc({ firstName: 'kyosuke', lastName: 'fujiki' }));
// console.log(testFunc({ firstName: 'kyosuke' }));
// let user: any = { firstName: 'kyosuke' }
// 以下の行のコードはいずれもコンパイラエラーが起こりません
// user.hello()
// user()
// user.age = 100
// user = 'hello'
// const n: number = user

// const sayHello = (name: string, greeting: string = 'Goodbye'): string => {
  // return `${greeting} ${name}`;
// };
// console.log(sayHello('kyosuke'));
// console.log(sayHello('kyosuke', 'Hello'));

// function sayHello(name: string): string {
//   return `Hello ${name}`
// }

// const toArray = (str: string): string[] => {
//   return str.split(',');
// }
// const sayAllItem = (toArrayFunc: (str: string) => string[]): void[] => {
//   return toArrayFunc('a,b,c,d,e').map((item) => {
// 	console.log(item)
//   })
// }
// sayAllItem(toArray)

// const age = 10
// console.log(age.length) // エラー: ageはnumber型なのでlengthプロパティはありません

// const user = {
//   name: 'Takuya',
//   age: 36
// }
// console.log(user.age.length) // エラー: ageはnumber型なのでlengthプロパティはありません

// function getUser() {
//   return {
//     name: 'Takuya',
//     age: 36
//   }
// }
// const user = getUser()
// console.log(user.age.length) // エラー: ageはnumber型なのでlengthプロパティはありません

// const names = ['Takuya', 'Yoshiki', 'Taketo']
// names.forEach((name) => {
//   // string型として扱われるので、関数名を間違えている呼び出しはコンパイル時エラーになります // 本来は toUpperCase が正しいです
//   console.log(name.toUppercase())
// })

// type Name = string;
// const test: Name = 'kyosuke';
// console.log(test);

// type Point = {
//   x: number;
//   y: number;
// };
// const printPoint = (point: Point) => {
//   console.log(`x: ${point.x}`);
//   console.log(`y: ${point.y}`);
// };
// printPoint({x: 100, y: 100});
// printPoint({t: 100, d: 100});

// type Formatter = (a: string) => string
// function printName(firstName: string, formatter: Formatter) {
//   console.log(formatter(firstName))
// }

// type Label = {
//   [key: string] : string
// }
// const labels: Label = {
//   topTitle: 'トップページのタイトルです',
//   topSubTitle: 'トップページのサブタイトルです',
//   topFeature1: 'トップページの機能1です',
//   topFeature2: 'トップページの機能2です'
// }
// // 値部分の型が合わないためエラー
// const hoge: Label = {
//   11: 'aa'
// }

// interface Point {
//   x: number;
//   y: number;
// }
// function printPoint(point: Point) {
//   console.log(`x座標は${point.x}です`)
//   console.log(`y座標は${point.y}です`)
//   console.log(`z座標は${point.z}です`)
// }
// interface Point {
//   z: number;
// }
// // 引数のオブジェクトにzが存在しないためコンパイル時にエラーになります
// // printPoint({ x: 100, y: 100})
// // 問題なく動作します
// printPoint({ x: 100, y: 100, z: 200})

// interface Point {
//   x: number;
//   y: number;
//   z?: number;
// }
// class MyPointTwoDemensions implements Point {
//   x: number;
//   y: number;
// }
// class MyPointThreeDemensions implements Point {
//   x: number;
//   y: number;
//   z: number;
// }

// class Point {
//   x: number;
//   y: number;
//   // 引数がない場合の初期値を指定します
//   constructor(x: number = 0, y: number = 0) {
//     this.x = x
//     this.y = y
//   }
//   // 戻り値がない関数を定義するためにvoidを指定します
//   moveX(n: number): void {
//     this.x += n
//   }
//   moveY(n: number): void {
//     this.y += n
//   }
// }
// const point = new Point()
// point.moveX(10)
// console.log(`${point.x}, ${point.y}`) // 10, 0

// class Point3D extends Point {
//   z: number;
//   constructor(x: number = 0, y: number = 0, z: number = 0) { // 継承元のコンストラクタを呼び出す
//     super(x, y)
//     this.z = z
//   }
//   moveZ(n: number): void {
//     this.z += n
//   }
// }
// const point3D = new Point3D()
// // 継承元のメソッドを呼び出すことができます
// point3D.moveX(10)
// point3D.moveZ(20)
// console.log(`${point3D.x}, ${point3D.y}, ${point3D.z}`) // 10, 0, 20


// // 頭のIはインタフェースであることを示すためのもの
// interface IUser {
//   name: string;
//   age: number;
//   sayHello: () => string; // 引数なしで文字列を返す
// }
// class User implements IUser {
//   name: string;
//   age: number;
//   constructor() {
//     this.name = ''
//     this.age = 0
//   }
// // インタフェースに定義されているメソッドを実装しない場合、コンパイル時エラーになります
//   sayHello(): string {
//     return `こんにちは、私は${this.name}、${this.age}歳です。` }
//   }
// const user = new User()
// user.name = 'kyosuke'
// user.age = 24
// console.log(user.sayHello()) // 'こんにちは、私はkyosuke、24歳です。'

// class BasePoint3D {
//   public x: number;
//   private y: number;
//   protected z: number;
// }

// // インスタンス化を行った場合のアクセス制御の例です
// const basePoint = new BasePoint3D()
// basePoint.x // OK
// basePoint.y // コンパイル時エラーが起きます。privateであるためアクセスできません
// basePoint.z // コンパイル時エラーが起きます。protectedもアクセスできません

// // クラスを継承した際のアクセス制御
// class ChildPoint extends BasePoint3D {
//   constructor() {
//     super()
//     this.x // OK
//     this.y // コンパイル時エラーが起きます。privateであるためアクセスできません
//     this.z // protectedは問題なくアクセスできます
//   }
// }

// jsでは下記
// const Direction = {
//   'Up': 0,
//   'Down': 1,
//   'Left': 2,
//   'Right': 3
// }
// enum Direction {
//   Up,
//   Down,
//   Left,
//   Right
// }
// // enum Directionを参照
// let direction: Direction = Direction.Left
// // 2という数字が出力されます
// console.log(direction)

// // enumを代入した変数に別の型の値を代入しようとするとエラーになります
// // direction = 'Left' // stringを入れようとするとエラー

// enum Direction {
//   Up = 'UP',
//   Down = 'DOWN',
//   Left = 'LEFT',
//   Right = 'RIGHT'
// }
// // たとえばAPIのパラメータとして文字列が渡されたケースを想定します
// const value = 'DOWN'
// // 文字列からEnum型に変換します
// // const enumValue = value as Direction
// // if (enumValue === Direction.Down) {
//   // console.log('Down is selected')
// // }
// if (value === Direction.Down) {
//   console.log('Down is selected')
// }

// // number型
// const test1 = (arg: number): number => {
//   return arg;
// }
// // string型
// const test2 = (arg: string): string => {
//   return arg;
// }
// test1(1); //=> 1
// test2("文字列"); //=> 文字列

// const test = <T>(arg: T): T => {
//   return arg;
// }
// test<number>(1); //=> 1
// test<string>("文字列"); //=> 文字列
// //※ Genericsでも型推論ができるので、引数から型が明示的にわかる場合は省略が可能
// console.log(test("文字列２")); //=> "文字列２"


// function testtest<T, U, P>(arg1:T, arg2: U, arg3: P): P {
//   return arg3;
// }
// console.log(testtest<string, boolean, number>("文字列", true, 4)); //=> 4


// class Klass<T> {
//   item: T;
//   constructor(item: T) {
//     this.item = item;
//   }
//   getItem(): T {
//     return this.item;
//   }
// }

// let strObj = new Klass<string>("文字列１");
// console.log(strObj.getItem()); //=> "文字列１"
// let numObj = new Klass<number>(5);
// console.log(numObj.getItem()); //=> 5

// interface IObj<T, U> {
//   item1: T;
//   item2: U;
// }
// let obj: IObj<string, number> = { item1: "文字列", item2: 2 } //= {key: "文字列", value: 2}

// const printId = (id: number | string) => {
//   console.log(id)
// }
// printId('kyosuke')
// printId(24)

// type Id = number | string;
// const printId = (id: Id) => {
//   console.log(id)
// }

// type Identity = {
//   id: number | string;
//   name: string;
// }
// type Contact = {
//   name: string;
//   email: string;
//   phone: string;
// }
// type IdentityOrContact = Identity | Contact

// // OK
// const id: IdentityOrContact = {
//   id: '111',
//   name: 'kyosuke'
// }
// // OK
// const contact: IdentityOrContact = {
//   name: 'kyosuke' ,
//   email: 'test@example.com',
//   phone: '012345678'
// }

// type Employee = Identity & Contact
// // OK
// const employee: Employee = {
//   id: '111',
//   name: 'Takuya',
//   email: 'test@example.com',
//   phone: '012345678'
// }
// // エラー: Contact情報のみでの変数定義はできません。idが必要です
// const employeeContact: Employee = {
//   name: 'Takuya',
//   email: 'test@example.com',
//   phone: '012345678'
// }

// エラーが常に返るような関数で決して値が正常に返らない場合にnever型を指定します
// function error(message: string): never {
//   throw new Error(message)
// }

// function foo(x: string | number | number[]): boolean {
//   if (typeof x === 'string') {
//     return true
//   } else if (typeof x === 'number') {
//     return false
//   }
//   // neverを利用することで明示的に値が返らないことをコンパイラに伝えることができます
//   // neverを使用しないとTypeScriptはコンパイルエラーになります
//   return error('Never happens')
// }

// foo('a')
// foo(1)
// foo([1])

// // 将来的にも定数が追加されるの可能性のあるenum型を定義します
// enum PageType {
//   ViewProfile,
//   EditProfile,
//   ChangePassword,
// }
// const getTitleText = (type: PageType) => {
//   switch (type) {
//     case PageType.ViewProfile:
//       return 'Setting'
//     case PageType.EditProfile:
//       return 'Edit Profile'
//     case PageType.ChangePassword:
//       return 'Change Password'
//     default:
//       // 決して起きないことをコンパイラに伝えるnever型に代入を行います
//       // これによって仮に将来PageTypeのenum型に定数が新規で追加された際に
//       // コンパイル時にエラーが起きるためバグを未然に防ぐ対応を行うことができます
//       const wrongType: never = type
//       throw new Error(`${wrongType} is not in PageType`)
//   }
// }

// nullになり得るsocialというプロパティの型を定義します
// interface User {
//   name: string
//   social?: {
//     facebook: boolean
//     twitter: boolean
//   }
// }
// let user: User
// user = { name: 'Takuya', social: { facebook: true, twitter: true } }
// console.log(user.social?.facebook) // true
// user = { name: 'Takuya' }
// // socialが存在しないケースでも以下のコードは実行時エラーになりません
// console.log(user.social?.facebook)
// // js
// if (user.social && user.social.twitter) {
//   console.log(user.social.twitter);
// }
//   // console.log(user.social.twitter);

// interface User {
//   name: string;
//   age: number;
//   email: string;
// }
// type UserKey = keyof User // 'name' | 'age' | 'email' というUnion型
// const key1: UserKey = 'name' // 代入可能
// const key2: UserKey = 'phone' // コンパイル時エラー

// function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
//   return obj[key]
// }
// const user: User = {
//   name: 'kyosuke',
//   age: 24,
//   email: 'test@example.com'
// }
// // nameは型のキーに存在するため正しくstring型の値が返ります
// const userName = getProperty(user, 'name')
// console.log(userName)
// genderはオブジェクトのキーに存在しないため、コンパイル時エラーになります
// const userGender = getProperty(user, 'gender')

// type User = {
//   name: string;
//   gender: string;
// }
// type UserReadonly = Readonly<User>

// let user: User = { name: 'Takuya', gender: 'Male' }
// let userReadonly: UserReadonly = { name: 'Takuya', gender: 'Male' }
// user.name = 'Yoshiki' // OK
// // userReadonly.name = 'Yoshiki' // コンパイル時エラー


// const x: unknown = 123
// const y: unknown = 'Hello'

// console.log(x.toFixed(1)) // error
// console.log(y.toLowerCase()) // error

// if (typeof x === 'number') {
//   console.log(x.toFixed(1)) // 123.0
// }
// if (typeof y === 'string') {
//   console.log(y.toLowerCase()) // hello
// }

// 非同期関数の定義します
// function fetchFromServer(id: string): Promise<{success: boolean}> {
//   return new Promise(resolve => {
//     setTimeout(() => {
//       resolve({success: true})
//     }, 100)})
// }
// // 非同期処理を含むasync functionの戻り値の型はPromiseとなります
// async function asyncFunc(): Promise<string> {
//   // Promiseな値をawaitすると中身が取り出せる(ように見える)
//   const result = await fetchFromServer('111')
//   return `The result: ${result.success}`
// }
// // Promiseとして扱う際は以下のように記述します
// asyncFunc().then(result => console.log(result))


// const printName = (myName: string) => { 
//   console.log(`Hello ${myName}!`); 
// };
// const myName: number = 24; 

// // error TS2345: Argument of type 'number' is not assignable to parameter of type 'string'.
// printName(myName);

// const message = 'Hello!';
// // error TS2349: This expression is not callable. Type 'String' has no call signatures.
// message();


// var var_val: string = true;
// let let_val: number = '11';
// const const_val: boolean = 11;

// const array: string[] = ['str'];

// const array2: string[] = [...array, 'string'];
// console.log(array2)

// // 型情報と要素の方があっていないためエラー
// const array3: number[] = [11, 24];
// console.log(array3)

// const mixedArrayU: (string|number|boolean)[] = [1, 1, 1];
// const mixedArrayT: [boolean, string, number] = [false, 'foo', 1];

// const userObjects: { name: string, age: number, address?: string } = {
//   name: 'kyosuke',
//   age: 24,
// };

// const testFunc = (obj: { firstName: string, lastName?: string }) => {
//   return `${obj.firstName} ${obj.lastName}`;
// };
// console.log(testFunc({ firstName: 'kyosuke', lastName: 'fujiki' }));
// console.log(testFunc({ firstName: 'kyosuke' })); // 問題なく実行される`


// let user: {firstName: string} = { firstName: 'kyosuke' };
//  // 以下の行のコードはいずれもコンパイラエラーが起こりません
// user.hello();
// user();
// user.age = 100;
// user = 'hello';
// const n: number = user;

// (name: string) => string
// const sayHello = (name: string): string => {
//   return `Hello ${name}`;
//   // return 11;
// };
// sayHello('kyosuke');


// const toArray = (str: string): string[] => {
//   return str.split(',');
// };


// // (引数名: 引数の型) => 戻り値の型
// const sayAllItem = (toArrayFunc: (str: string) => string[]): void[] => {
//   return toArrayFunc('a,b,c,d,e').map((item) => {
//     console.log(item);
//   });
// };

// // [a, b, c, d, e]


// // sayAllItem(toArray);
// // sayAllItem('kyosuke'); // 型が違うためエラー

// const age = 10;
// console.log(age.length); // エラー: ageはnumber型なのでlengthプロパティはありません

// const user = {
//   name: 'Takuya',
//   age: 36
// };
// console.log(user.age.length); // エラー: ageはnumber型なのでlengthプロパティはありません

// function getUser() {
//   return {
//     name: 'Takuya',
//     age: 36
//   };
// }
// const user = getUser();
// console.log(user.age.length); // エラー: ageはnumber型なのでlengthプロパティはありません

// const names = ['Takuya', 'Yoshiki', 'Taketo'];
// names.forEach((name) => {
//   // string型として扱われるので、関数名を間違えている呼び出しはコンパイル時エラーになります
//   // 本来は toUpperCase が正しい
//   console.log(name.toUppercase());
// })

// // window.confirm関数の返り型はbooelanを返すことをTypeScriptは知っているため
// // 代入する関数の型が一致しない場合コンパイル時エラーになります
// window.confirm = () => {
//   // booleanをreturnしない限りエラーになります
//   console.log('confirm関数');
// }

// const myCanvas = document.getElementById('main_canvas')
// // error TS2339: Property 'width' does not exist on type 'HTMLElement'.
// console.log(myCanvas.width)

// const test = {} as string;
// console.log(test)

// 変数
// type Name = string;
// const test: Name = 'kyosuke';
// console.log(test);

// // オブジェクト
// type Point = {
//   x: number;
//   y: number;
// };
// const printPoint = (point: Point) => {
//   console.log(`x: ${point.x}`);
//   console.log(`y: ${point.y}`);
// };
// printPoint({x: 100, y: 100});

// // 関数
// type Formatter = (a: string) => string;
// function printName(firstName: string, formatter: Formatter) {
//   console.log(formatter(firstName));
// }

// interface Point {
//   x: number;
//   y: number;
// }
// function printPoint(point: Point) {
//   console.log(`x座標は${point.x}です`)
//   console.log(`y座標は${point.y}です`)
//   console.log(`z座標は${point.z}です`)
// }
// interface Point {
//   z: number;
// }
// 引数のオブジェクトにzが存在しないためコンパイル時にエラーになります
// printPoint({ x: 100, y: 100});
// 問題なく動作します
// printPoint({ x: 100, y: 100, z: 200});

// interface Point {
//   x: number;
//   y: number;
//   z?: number; // オプショナルなプロパティにする
// }
// class MyPointTwoDemensions implements Point {
//   x: number;
//   y: number;
// }
// class MyPointThreeDemensions implements Point {
//   x: number;
//   y: number;
//   z: number;
// }

// // 頭のIはインタフェースであることを示すためのもの
// interface IUser {
//   name: string;
//   age: number;
//   sayHello: () => string; // 引数なしで文字列を返す
// }
// class User implements IUser {
//   name: string;
//   age: number;
//   constructor() {
//     this.name = '';
//     this.age = 0;
//   }
//   // インタフェースに定義されているメソッドを実装しない場合、コンパイル時エラーになります
//   sayHello(): string {
//     return `こんにちは、私は${this.name}、${this.age}歳です。`;
//   }
// }
// const user = new User();
// user.name = 'kyosuke';
// user.age = 24;
// console.log(user.sayHello()); // 'こんにちは、私はkyosuke、24歳です。'

// class BasePoint3D {
//   public x: number;
//   private y: number;
//   protected z: number;
// }
// // インスタンス化を行った場合のアクセス制御の例です
// const basePoint = new BasePoint3D();
// console.log(basePoint.x); // OK
// console.log(basePoint.y); // コンパイル時エラーが起きます。privateであるためアクセスできません
// console.log(basePoint.z); // コンパイル時エラーが起きます。protectedもアクセスできません
// // クラスを継承した際のアクセス制御
// class ChildPoint extends BasePoint3D {
//   constructor() {
//     super();
//     this.x; // OK
//     this.y; // コンパイル時エラーが起きます。privateであるためアクセスできません
//     this.z; // protectedは問題なくアクセスできます
//   }
// }

// jsでは下記
// const Direction = {
//   'Up': 0,
//   'Down': 1,
//   'Left': 2,
//   'Right': 3
// }
// enum numDirection {
//   Up,
//   Down,
//   Left,
//   Right
// }

// // enum Directionを参照
// let direction: numDirection = numDirection.Left;
// // 2という数字が出力されます
// console.log(direction);
// if (direction === 2) {
//   console.log('Left is selected')
// }

// // // enumを代入した変数に別の型の値を代入しようとするとエラーになります
// // direction = 'Left'; // stringを入れようとするとエラー

// enum Direction {
//   Up = 'UP',
//   Down = 'DOWN',
//   Left = 'LEFT',
//   Right = 'RIGHT'
// }
// // たとえばAPIのパラメータとして文字列が渡されたケースを想定します
// const value = 'DOWN';
// // 文字列からEnum型に変換します
// const enumValue = value as Direction;
// console.log(typeof enumValue)
// // if (value === Direction.Down) {
// //   console.log('Down is selected');
// // }

// // number型
// const test1 = (arg: number): number => {
//   return arg;
// };
// // string型
// const test2 = (arg: string): string => {
//   return arg;
// };
// test1(1); //=> 1
// test2("文字列"); //=> 文字列

// ---------------------------------
// const test = <T>(arg: T): T => {
//   return arg;
// };
// test<number>(1); //=> 1
// test<string>("文字列"); //=> 文字列

// // console.log(test<boolean>());

// // 複数の引数を取ることも可能
// function testTest<T, U, P>(arg1:T, arg2: U, arg3: P): U|T|P {
//   const tmp: T = arg1
//   return tmp
// }
// console.log(testTest<string, string, number>("文字列", 'false', 4));

// インラインで指定
// const printId = (id: number | string) => {
//   console.log(id);
// };
// printId('kyosuke');
// printId(24);

// // 型エイリアスで指定
// type Id = number | string;
// const await test = () => {
//   const id = async fetch('url...')
//   return id
// }
// type tmpId = string | undefined;
// const jsonId : tmpId = test()


// const printId = (id: Id) => {
//   console.log(id);
// };

// // 型エイリアスを和集合
// type Identity = {
//   id: number | string;
//   name: string;
// };
// type Contact = {
//   name: string;
//   email: string;
//   phone: string;
// };
// type IdentityOrContact = Identity | Contact;
// // OK
// const id: IdentityOrContact = {
//   id: '111',
//   name: 'kyosuke'
// };
// // OK
// const contact: IdentityOrContact = {
//   name: 'kyosuke',
//   email: 'test@example.com',
//   phone: '012345678'
// };

// type Employee = Identity & Contact;
// // OK
// const employee: Employee = {
//   id: '111',
//   name: 'Takuya',
//   email: 'test@example.com',
//   phone: '012345678'
// };
// // エラー: idがないのでエラー
// const employeeContact: Employee = {
//   id: '112',
//   name: 'Takuya',
//   email: 'test@example.com',
//   phone: '012345678'
// };

// nullになり得るsocialというプロパティの型を定義します
// interface User {
//   name: string
//   social?: {
//     facebook: boolean
//     twitter: boolean
//   }
// }
// let user: User
// user = { name: 'Takuya', social: { facebook: true, twitter: true } }
// console.log(user.social?.facebook) // true
// user = { name: 'Takuya' }
// // socialが存在しないケースでも以下のコードは実行時エラーになりません
// console.log(user.social?.facebook)

// // js
// if (user.social && user.social.twitter) {
//   console.log(user.social.twitter);
// }

// userがnullの場合、実行時エラーになる可能性があるプロパティへのアクセスはコンパイルエラー
// !を用いて明示的に指定することでコンパイルエラーを抑制
// function processUser(user?: User) {
//   let s = user!.name
//   console.log(s)
// }

// const foo = (arg: number | string | boolean) => {
//   if (typeof arg === "number") {
//     console.log(typeof arg)
//     const y = arg; // const y: number
//   } else {
//     console.log(typeof arg)
//     const z = arg;
//   }
// };

// foo(1)
// foo('kyosuke')

// interface User {
//   name: string;
//   age: number;
//   email: string;
// }
// type UserKey = keyof User // 'name' | 'age' | 'email' というUnion型
// const key1: UserKey/*('name | 'age' | 'email')*/ = 'name' // 代入可能
// // const key2: UserKey = 'phone' // コンパイル時エラー

// // 第1引数に渡したオブジェクトの型のプロパティ名のUnion型と、第2引数で渡す値の型が一致しない場合型エラーになります
// // T[K]によりキーに対応する型が戻り値の型となります
// // (たとえば上記Userのageをkeyに渡した場合、戻り値の方はnumberになります)
// function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
//   return obj[key]
// }
// const user: User = {
//   name: 'kyosuke',
//   age: 24,
//   email: 'test@example.com'
// }
// // nameは型のキーに存在するため正しくstring型の値が返る
// const userName = getProperty<User, keyof User/*(name | age | email)*/>(user, 'name')
// console.log(userName)
// genderはオブジェクトのキーに存在しないため、コンパイル時エラー
// const userGender = getProperty<User, keyof User>(user, 'gender')

// type User = {
//   readonly name: string;
//   readonly gender: string;
// }
// type TmpUser = {
//   name: string;
//   gender: string;
// }

// type readonlyUser = Readonly<TmpUser>
// // let user: User = { name: 'kyosuke', gender: 'Male' }
// const tmpUser: readonlyUser = { name: 'kyosuke', gender: 'Male' }
// tmpUser.gender = 'Female' // コンパイル時エラー
// console.log(tmpUser)

const x: unknown = 123
const y: unknown = 'Hello'

console.log(x.toFixed(1)) // error
console.log(y.toLowerCase()) // error

if (typeof x === 'number') {
  console.log(x.toFixed(1)) // 123.0
}
if (typeof y === 'string') {
  console.log(y.toLowerCase()) // hello
}
