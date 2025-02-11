package jsonpath_test

import (
	"encoding/json"
	"testing"

	. "github.com/vnay92/jsonpath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Jpath", func() {
	Context("Simple key value", func() {
		It("generates json", func() {
			in := map[string]string{"key": "value"}
			actual, err := Marshal(in, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal([]byte(`{"key":"value"}`)))
		})
	})
	Context("Simple embeddens key value", func() {
		It("generates json", func() {
			in := map[string]string{"price.value": "100.00"}
			actual, err := Marshal(in, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal([]byte(`{"price":{"value":"100.00"}}`)))
		})
	})
	Context("Panic atack with number and dot", func() {
		It("generates json skipping the wrong keys - first hashes", func() {
			in := map[string]string{"one": "1233", "2. subcategory": "booooom", "two": "2"}
			_, err := Marshal(in, "")
			Expect(err).To(HaveOccurred())
		})
		It("generates json skipping the wrong keys - first arrays", func() {
			in := map[string]string{"2. subcategory": "booooom", "one": "1233", "two": "2"}
			_, err := Marshal(in, "")
			Expect(err).To(HaveOccurred())
		})
	})
	Context("Long embeddens key value", func() {
		It("generates json", func() {
			in := map[string]string{"price.value1.value2": "100.00"}
			actual, err := Marshal(in, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal([]byte(`{"price":{"value1":{"value2":"100.00"}}}`)))
		})
	})
	Context("Simple embeddens few key values", func() {
		It("generates json", func() {
			in := map[string]string{"price.value": "100.00", "price.currency": "EU"}
			actual, err := Marshal(in, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal([]byte(`{"price":{"currency":"EU","value":"100.00"}}`)))
		})
	})
	Context("Simple embeddens few key values different levels", func() {
		It("generates json", func() {
			in := map[string]string{"price.value": "100.00", "price.currency": "EU", "shipping.value": "99.00", "shipping.currency": "UA"}
			actual, err := Marshal(in, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal([]byte(`{"price":{"currency":"EU","value":"100.00"},"shipping":{"currency":"UA","value":"99.00"}}`)))
		})
	})
	Context("Simple embeddens few key values and array with one value", func() {
		It("generates json", func() {
			in := map[string]string{"prices.0": "100.00"}
			actual, err := Marshal(in, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal([]byte(`{"prices":["100.00"]}`)))
		})
	})
	Context("Simple embeddens few key values and array with two values", func() {
		It("generates json", func() {
			in := map[string]string{"prices.1": "100.00", "prices.0": "10.00"}
			actual, err := Marshal(in, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal([]byte(`{"prices":["10.00","100.00"]}`)))
		})
	})
	Context("Simple embeddens few key values and array with three values", func() {
		It("generates json", func() {
			in := map[string]string{"prices.2": "100.00", "prices.0": "10.00"}
			actual, err := Marshal(in, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal([]byte(`{"prices":["10.00",null,"100.00"]}`)))
		})
	})
	Context("Simple embeddens few key values and array with shipping", func() {
		It("generates json", func() {
			in := map[string]string{"price.value": "100.00", "price.currency": "EU", "shipping.0.country": "GB", "shipping.0.service": "Standart shipping", "shipping.0.price.value": "33", "shipping.0.price.curency": "GBP"}
			actual, err := Marshal(in, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal([]byte(`{"price":{"currency":"EU","value":"100.00"},"shipping":[{"country":"GB","price":{"curency":"GBP","value":"33"},"service":"Standart shipping"}]}`)))
		})
	})
	Context("Arrays", func() {
		Context("key value", func() {
			It("generates json", func() {
				in := map[string]string{"0.value": "100.00"}
				actual, err := Marshal(in, "")
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal([]byte(`[{"value":"100.00"}]`)))
			})
		})
		Context("array key", func() {
			It("generates array", func() {
				in := map[string]string{"0.value.[]": "1,2,3,4"}
				actual, err := Marshal(in, "")
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal([]byte(`[{"value":["1","2","3","4"]}]`)))
			})
		})
		Context("key value with num", func() {
			It("generates json", func() {
				in := map[string]string{"0.value.num()": "100.00"}
				actual, err := Marshal(in, "")
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal([]byte(`[{"value":100}]`)))
			})
		})
		Context("key value with float", func() {
			It("generates json", func() {
				in := map[string]string{"0.value.num()": "100.12"}
				actual, err := Marshal(in, "")
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal([]byte(`[{"value":100.12}]`)))
			})
		})
		Context("key value with bool", func() {
			It("generates json for true", func() {
				in := map[string]string{"0.value.bool()": "true"}
				actual, err := Marshal(in, "")
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal([]byte(`[{"value":true}]`)))
			})
			It("generates json for false", func() {
				in := map[string]string{"0.value.bool()": "false"}
				actual, err := Marshal(in, "")
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal([]byte(`[{"value":false}]`)))
			})
			It("generates json for empty", func() {
				in := map[string]string{"0.value.bool()": ""}
				actual, err := Marshal(in, "")
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal([]byte(`[{"value":false}]`)))
			})
			It("generates json for not bool value", func() {
				in := map[string]string{"0.value.bool()": "1234"}
				actual, err := Marshal(in, "")
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal([]byte(`[{"value":false}]`)))
			})
		})
	})
	Measure("it should do something hard efficiently", func(b Benchmarker) {
		runtime := b.Time("runtime", func() {
			in := map[string]string{"price.value": "100.00", "price.currency": "EU", "shipping.0.country": "GB", "shipping.0.service": "Standart shipping", "shipping.0.price.value": "33", "shipping.0.price.curency": "GBP"}
			actual, err := Marshal(in, "")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal([]byte(`{"price":{"currency":"EU","value":"100.00"},"shipping":[{"country":"GB","price":{"curency":"GBP","value":"33"},"service":"Standart shipping"}]}`)))
		})
		Expect(runtime.Seconds()).Should(BeNumerically("<", 0.2), "SomethingHard() shouldn't take too long.")
	}, 10)
})

func BenchmarkComplexJSONPathArray(b *testing.B) {
	in := map[string]string{"price.value": "100.00", "price.currency": "EU", "shipping.0.country": "GB", "shipping.0.service": "Standart shipping", "shipping.0.price.value": "33", "shipping.0.price.curency": "GBP"}
	for n := 0; n < b.N; n++ {
		Marshal(in, "")
	}
}
func BenchmarkSimpleJSONPathArrayWithNum(b *testing.B) {
	in := map[string]string{"0.value.num()": "100.12"}
	for n := 0; n < b.N; n++ {
		Marshal(in, "")
	}
}
func BenchmarkSimpleJSONPathArrayWithBool(b *testing.B) {
	in := map[string]string{"0.value.bool()": "true"}
	for n := 0; n < b.N; n++ {
		Marshal(in, "")
	}
}

func BenchmarkSimpleJSONPathArrayInsideArray(b *testing.B) {
	in := map[string]string{"0.value.[]": "1,2,3,4,5,6"}
	for n := 0; n < b.N; n++ {
		Marshal(in, "")
	}
}

func BenchmarkSimpleJSONPathArrays(b *testing.B) {
	in := map[string]string{"value.[]": "1,2,3,4,5,6"}
	for n := 0; n < b.N; n++ {
		Marshal(in, "")
	}
}

func BenchmarkSimpleJSONPathSimple(b *testing.B) {
	in := map[string]string{"value": "100.12"}
	for n := 0; n < b.N; n++ {
		Marshal(in, "")
	}
}
func BenchmarkJSONNative(b *testing.B) {
	in := map[string]string{"0.value.num()": "100.12"}
	for n := 0; n < b.N; n++ {
		json.Marshal(in)
	}
}
