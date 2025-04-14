package ui

import (
	"fmt"
	"golang.design/x/hotkey"
	"strings"
)

// parseHotKey 把 "Ctrl+Shift+T" 之类的字符串拆成 修饰键切片 + 主键
func ParseHotKey(combo string) (mods []hotkey.Modifier, key hotkey.Key, err error) {
	parts := strings.Split(combo, "+")
	if len(parts) < 2 {
		return nil, 0, fmt.Errorf("必须至少包含一个修饰键和一个主键，用 + 分隔")
	}
	// 最后一个当主键，其余都是修饰键
	rawMods := parts[:len(parts)-1]
	rawKey := parts[len(parts)-1]

	// 解析修饰键
	for _, m := range rawMods {
		switch strings.ToLower(strings.TrimSpace(m)) {
		case "ctrl":
			mods = append(mods, hotkey.ModCtrl)
		case "shift":
			mods = append(mods, hotkey.ModShift)
		case "alt":
			mods = append(mods, hotkey.ModAlt)
		default:
			return nil, 0, fmt.Errorf("不支持的修饰键：%s", m)
		}
	}
	// 解析主键（这里只举几个例子，按需补全）
	switch strings.ToUpper(strings.TrimSpace(rawKey)) {
	case "A":
		key = hotkey.KeyA
	case "B":
		key = hotkey.KeyB
	case "C":
		key = hotkey.KeyC
	case "D":
		key = hotkey.KeyD
	case "E":
		key = hotkey.KeyE
	case "F":
		key = hotkey.KeyF
	case "G":
		key = hotkey.KeyG
	case "H":
		key = hotkey.KeyH
	case "I":
		key = hotkey.KeyI
	case "J":
		key = hotkey.KeyJ
	case "K":
		key = hotkey.KeyK
	case "L":
		key = hotkey.KeyL
	case "M":
		key = hotkey.KeyM
	case "N":
		key = hotkey.KeyN
	case "O":
		key = hotkey.KeyO
	case "P":
		key = hotkey.KeyP
	case "Q":
		key = hotkey.KeyQ
	case "R":
		key = hotkey.KeyR
	case "S":
		key = hotkey.KeyS
	case "T":
		key = hotkey.KeyT
	case "U":
		key = hotkey.KeyU
	case "V":
		key = hotkey.KeyV
	case "W":
		key = hotkey.KeyW
	case "X":
		key = hotkey.KeyX
	case "Y":
		key = hotkey.KeyY
	case "Z":
		key = hotkey.KeyZ
	case "F1":
		key = hotkey.KeyF1
	case "F2":
		key = hotkey.KeyF2
	case "F3":
		key = hotkey.KeyF3
	case "F4":
		key = hotkey.KeyF4
	case "F5":
		key = hotkey.KeyF5
	case "F6":
		key = hotkey.KeyF6
	case "F7":
		key = hotkey.KeyF7
	case "F8":
		key = hotkey.KeyF8
	case "F9":
		key = hotkey.KeyF9
	case "F10":
		key = hotkey.KeyF10
	case "F11":
		key = hotkey.KeyF11
	default:
		return nil, 0, fmt.Errorf("不支持的主键：%s", rawKey)
	}
	return mods, key, nil
}
