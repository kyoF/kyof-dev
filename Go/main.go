package main

import (
    "fmt"
    "sort"
)

type Entry struct {
    Name string
    Value int
}

type List []Entry

func (l List) Len() int {
    return len(l)
}

func (l List) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
    if l[i].Value == l[j].Value {
        return (l[i].Name < l[j].Name)
    } else {
        return (l[i].Value < l[j].Value)
    }
}

func main() {
    m := map[string]int{"S":1, "J":4, "A":3, "N":3}
    lt := List{}
    for k, v := range m {
        e := Entry{k, v}
        lt = append(lt, e)
    }
    sort.Sort(lt)
    fmt.Println(lt)
    fmt.Println("-----")

    sort.Sort(sort.Reverse(lt))
    fmt.Println(lt)

    el := []Entry {
        {"A",20},
        {"C",30},
        {"c",2},
        {"d",1},
        {"B",320},
        {"Aa",80},
        {"AA",220},
        {"X",80},
        {"D",830},
        {"F",2},
    }
    fmt.Println(el)

    sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })
    sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })
    fmt.Println("-----")
    fmt.Println(el)

    sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })
    sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })
    fmt.Println("-----")
    fmt.Println(el)
}
