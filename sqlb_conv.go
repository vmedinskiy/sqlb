package sqlb

func convInt64Slice(in []int64) any {
	res := make(ArrayAppender[Int64A], len(in))
	for i, v := range in {
		res[i] = Int64A(v)
	}
	return StringSQA(res.String())
}

func convIntSlice(in []int) any {
	res := make(ArrayAppender[Int64A], len(in))
	for i, v := range in {
		res[i] = Int64A(v)
	}
	return StringSQA(res.String())
}

func convStringSlice(in []string) any {
	res := make(ArrayAppender[StringDQA], len(in))
	for i, v := range in {
		res[i] = StringDQA(v)
	}
	return StringSQA(res.String())
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
			res[k] = StringSQA(v)
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
