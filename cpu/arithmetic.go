package cpu

import "go6502/mem"

func (C *CPU) op_ADC_IM(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ADC_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ADC_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ADC_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ADC_ABX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ADC_ABY(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ADC_INX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_ADC_INY(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_SBC_IM(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_SBC_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_SBC_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABY(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_INX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_INY(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_CMP_IM(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_CMP_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_CMP_ZPX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_CMP_ABS(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_CMP_ABX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_CMP_ABY(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_CMP_INX(mem *mem.Memory) { C.opName = "ToDO" }
func (C *CPU) op_CMP_INY(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_CPX_IM(mem *mem.Memory) {
	C.opName = "CPX Im"
	value := C.fetchByte(mem)
	C.setC(C.X, value)
	res := C.X - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CPX_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_CPX_ABS(mem *mem.Memory) { C.opName = "ToDO" }

func (C *CPU) op_CPY_IM(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_CPY_ZP(mem *mem.Memory)  { C.opName = "ToDO" }
func (C *CPU) op_CPY_ABS(mem *mem.Memory) { C.opName = "ToDO" }
