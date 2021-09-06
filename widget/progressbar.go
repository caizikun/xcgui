package widget

import (
	xc "github.com/twgh/xcgui"
)

type ProgressBar struct {
	Element

	HEle int
}

// 进度条_创建, 创建进度条元素, 返回元素句柄
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口
func NewProgressBar(x int, y int, cx int, cy int, hParent int) *ProgressBar {
	p := &ProgressBar{
		HEle: xc.XProgBar_Create(x, y, cx, cy, hParent),
	}
	p.HEle_ = p.HEle
	return p
}

// 进度条_置范围, 设置范围
// range_: 范围.
func (p *ProgressBar) SetRange(range_ int) int {
	return xc.XProgBar_SetRange(p.HEle, range_)
}

// 进度条_取范围
func (p *ProgressBar) GetRange() int {
	return xc.XProgBar_GetRange(p.HEle)
}

// 进度条_置进度图片
// hImage: 图片句柄.
func (p *ProgressBar) SetImageLoad(hImage int) int {
	return xc.XProgBar_SetImageLoad(p.HEle, hImage)
}

// 进度条_置两端大小, 设置两端间隔大小
// leftSize: 左边间隔大小.
// rightSize: 右边间隔大小.
func (p *ProgressBar) SetSpaceTwo(leftSize int, rightSize int) int {
	return xc.XProgBar_SetSpaceTwo(p.HEle, leftSize, rightSize)
}

// 进度条_置进度, 设置位置点
// pos: 位置点.
func (p *ProgressBar) SetPos(pos int) int {
	return xc.XProgBar_SetPos(p.HEle, pos)
}

// 进度条_取进度, 获取当前位置点
func (p *ProgressBar) GetPos() int {
	return xc.XProgBar_GetPos(p.HEle)
}

// 进度条_置水平, 设置水平或垂直
// bHorizon: 水平或垂直.
func (p *ProgressBar) SetHorizon(bHorizon bool) int {
	return xc.XProgBar_SetHorizon(p.HEle, bHorizon)
}