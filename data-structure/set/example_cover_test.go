package set

// 这个示例实现了简单的集合覆盖算法

import "fmt"

// cover 实现集合覆盖
// members是所有技能集合
// subsets是members的子集A1， 到An组成的集合。 cover 从subsets中找出覆盖members最高的项
func cover(members *Set, subsets *Set) (*Set, error) {
	members = members.Clone()
	subsets = subsets.Clone()
	rt := New(DefaultMatch)
	var maxLen int
	var maxMember *Set
	for members.Len() > 0 && subsets.Len() > 0 {
		maxLen = 0

		// 遍历subset
		subsets.Each(func(data interface{}, _ int) bool {
			member, ok := data.(*Set)
			if !ok {
				panic("Type error")
			}
			i := member.Intersection(members)
			if i.Len() > maxLen {
				maxMember = member
				maxLen = i.Len()
			}
			return false
		})

		// 没有任何子集覆盖
		if maxLen == 0 {
			return nil, ErrNotFound
		}
		rt.Insert(maxMember)
		members = members.Diff(maxMember)
		subsets.Remove(maxMember)
	}
	if rt.Len() == 0 {
		return nil, ErrNotFound
	}
	return rt, nil
}

func Example() {
	skills := New(DefaultMatch, "c++", "go", "python", "ruby", "java")
	group := New(func(a, b interface{}) bool {
		sa := a.(*Set)
		sb := b.(*Set)
		return sa.Equal(sb)
	},
		New(DefaultMatch, "python", "go", "ruby"),
		New(DefaultMatch, "c++", "python"),
		New(DefaultMatch, "java", "ruby"),
		New(DefaultMatch, "go", "java", "ruby"),
	)
	best, err := cover(skills, group)
	if err != nil {
		fmt.Println("覆盖失败")
	} else {
		fmt.Println(best)
	}
}
