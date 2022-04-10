package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("simple test", func(t *testing.T) {
		l := NewList()
		back := l.PushFront("back")   // [back]
		med1 := l.PushFront("med1")   // [med1 back]
		med2 := l.PushFront("med2")   // [med2 med1 back]
		front := l.PushFront("front") // [front med2 med1 back]

		l.Remove(med2) // [front med1 back]
		require.Equal(t, front, l.Front())
		require.Equal(t, back, l.Back())
		require.Equal(t, med1, l.Front().Prev)
		require.Equal(t, med1, l.Back().Next)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		// Ранее тут был Next,
		// но как у самого верхнего элемента может быть следующий?
		middle := l.Front().Prev // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		// Ранее тут был Next,
		// но как у самого верхнего элемента может быть следующий?
		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Prev {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
