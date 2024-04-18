package game

type Cell struct {
	X       int
	Y       int
	Content string
	Kind    string
}

func (cell *Cell) isEdgeTop() bool {
	return cell.X == 0
}

func (cell *Cell) isEdgeBottom() bool {
	return cell.X == 31
}

func (cell *Cell) isEdgeLeft() bool {
	return cell.Y == 0
}

func (cell *Cell) isEdgeRight() bool {
	return cell.Y == 31
}

func (cell *Cell) isCornerTopLeft() bool {
	if cell.isEdgeTop() && cell.isEdgeLeft() {
		return true
	}
	return false
}

func (cell *Cell) isCornerTopRight() bool {
	if cell.isEdgeTop() && cell.isEdgeRight() {
		return true
	}
	return false
}
func (cell *Cell) isCornerBottomLeft() bool {
	if cell.isEdgeBottom() && cell.isEdgeLeft() {
		return true
	}
	return false
}
func (cell *Cell) isCornerBottomRight() bool {
	if cell.isEdgeBottom() && cell.isEdgeRight() {
		return true
	}
	return false
}

func (cell *Cell) DetectContent() string {
	if cell.isEdgeTop() {
		if cell.isCornerTopLeft() {
			return Symbols.CornerTopLeft
		}
		if cell.isCornerTopRight() {
			return Symbols.CornerTopRight
		}
		return Symbols.EdgeTop
	}

	if cell.isEdgeBottom() {
		if cell.isCornerBottomLeft() {
			return Symbols.CornerBottomLeft
		}
		if cell.isCornerBottomRight() {
			return Symbols.CornerBottomRight
		}
		return Symbols.EdgeBottom
	}

	if cell.isEdgeLeft() {
		return Symbols.EdgeLeft
	}
	if cell.isEdgeRight() {
		return Symbols.EdgeRight
	}

	return Symbols.EmptyCell
}

func (cell *Cell) IsWall() bool {
	return cell.Kind == "wall"
}

func (cell *Cell) IsTower() bool {
	return cell.Kind == "tower"
}
