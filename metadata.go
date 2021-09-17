package bmmeta

// Metadata
type BasmMeta struct {
	metaData map[string]string
}

func (bm *BasmMeta) loopMeta() map[string]string {
	if bm != nil && bm.metaData != nil {
		return bm.metaData
	}
	return map[string]string{}
}

func (bm *BasmMeta) listMeta() string {
	result := ""
	if bm != nil && bm.metaData != nil {
		for key, val := range bm.metaData {
			result += "," + key + ":" + val
		}
	}
	if len(result) != 0 {
		return result[1:]
	}
	return result
}

func (bm *BasmMeta) getMeta(meta string) string {
	if bm != nil && bm.metaData != nil {
		if value, ok := bm.metaData[meta]; ok {
			return value
		} else {
			return ""
		}
	}
	return ""
}

func (bm *BasmMeta) rmMeta(meta string) {
	if bm != nil && bm.metaData != nil {
		if _, ok := bm.metaData[meta]; ok {
			delete(bm.metaData, meta)
		}
	}
}

func (bm *BasmMeta) setMeta(meta string, value string) *BasmMeta {
	if bm == nil {
		newbm := new(BasmMeta)
		newbm.metaData = make(map[string]string)
		newbm.metaData[meta] = value
		return newbm
	}
	bm.metaData[meta] = value
	return bm

}
