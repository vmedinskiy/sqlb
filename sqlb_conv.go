package sqlb

func convInt64Slice(in []int64) any {
	res := make(ArrayAppender[Int64A], len(in))
	for i, v := range in {
		res[i] = Int64A(v)
	}
	return StringQA(res.String())
}

func convIntSlice(in []int) any {
	res := make(ArrayAppender[Int64A], len(in))
	for i, v := range in {
		res[i] = Int64A(v)
	}
	return StringQA(res.String())
}

func convStringSlice(in []string) any {
	res := make(ArrayAppender[StringQA], len(in))
	for i, v := range in {
		res[i] = StringQA(v)
	}
	return StringQA(res.String())
}

func conv(in map[string]any) Dict {
	res := make(map[string]any, len(in))
	for k, v := range in {
		switch v := v.(type) {
		case int64:
			res[k] = v
		case float64:
			res[k] = v
		case string:
			res[k] = StringQA(v)
		case []int:
			res[k] = convIntSlice(v)
		case []int64:
			res[k] = convInt64Slice(v)
		case []string:
			res[k] = convStringSlice(v)
		default:
			res[k] = v
		}
	}
	return res
}
