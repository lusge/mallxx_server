package utils

const PAGE_ZSIZE = 2000

func Pager(pageNum, pageSize int32) (int32, int32) {
	//分页计算
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	if pageSize > PAGE_ZSIZE {
		pageSize = PAGE_ZSIZE
	}

	start := (pageNum - 1) * pageSize

	return start, pageSize
}
