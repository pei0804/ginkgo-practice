package calc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tikasan/ginkgo-practice/calc"
)

var _ = Describe("Calc", func() {
	var hoge calc.Hoge
	var err error

	Describe("データをロードする", func() {
		Context("正しいデータが格納された時", func() {
			BeforeEach(func() {
				hoge, err = calc.NewHoge(1, 2)
			})
			It("データを正しく取得する", func() {
				Expect(hoge.Atai1).To(Equal(1))
				Expect(hoge.Atai2).To(Equal(2))
			})
			It("エラーは発生してはならない", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("正しくデータが取得できない時", func() {
			BeforeEach(func() {
				hoge, err = calc.NewHoge(0, 2)
			})
			It("ゼロ値が返ってこなければいけない", func() {
				Expect(hoge).To(BeZero())
			})
			It("エラーが格納されていなければならない", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	BeforeEach(func() {
		hoge, err = calc.NewHoge(2, 1)
	})

	Describe("演算する", func() {
		It("足し算をして2+1=3が返ってくる", func() {
			Expect(hoge.Add()).To(Equal(3))
		})
		It("引き算をして2-1=1が返ってくる", func() {
			Expect(hoge.Sub()).To(Equal(1))
		})
	})
})
