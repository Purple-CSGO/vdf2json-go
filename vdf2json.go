package vdf2json_go

func ToJson(vdfData string) (string, error) {
	// - 按行遍历 如果有 \"\t\t\" -> \": \" 同时在结尾加上,
	// - 按行遍历 如果有 { → : { ... 如果有 } → },
	// - trim每行左边的 \t 和每行右边的换行符 \r 和 \n
	// - 遍历 如果有 },} → }}

	return "", nil
}
