package donut

import (
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
)

// View renders the current state of the model as a string.
func (m *Model) View() tea.View {
	var sb strings.Builder

	size := clamp(m.Size(), 32768, 155648)
	sb.Grow(size)

	if !m.mute {
		m.writeHeader(&sb)
	}

	hpad := (m.w - DonutW) / 2
	vpad := (m.h - DonutH) / 2

	for range vpad {
		sb.WriteByte('\n')
	}

	for i := range DonutH {
		for range hpad {
			sb.WriteByte(' ')
		}

		for j := range DonutW {
			if m.depth[i*DonutW+j] == 0 {
				sb.WriteByte(' ')
				continue
			}

			s := m.grid[i*DonutW+j]
			sb.WriteString(s.String())
		}
		sb.WriteByte('\n')
	}

	v := tea.NewView(sb.String())
	v.AltScreen = true

	return v
}

func (m *Model) writeHeader(sb *strings.Builder) {
	const (
		header0 = "fps [q]uit [m]ute [c]olor"
	)

	var buf [24]byte

	s := strconv.AppendFloat(buf[:0], m.FPS(), 'f', 0, 64)

	if len(s) < 2 {
		sb.WriteByte(' ')
	}

	sb.Write(s)
	sb.WriteString(header0)
	sb.WriteByte('\n')
}
