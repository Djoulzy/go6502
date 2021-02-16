package main

func (C *CPU) op_ADC_IM(mem *Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ADC_ZP(mem *Memory)  { C.opName = "ToDO" }
func (C *CPU) op_ADC_ZPX(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_ADC_ABS(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_ADC_ABX(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_ADC_ABY(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_ADC_INX(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_ADC_INY(mem *Memory) { C.opName = "ToDO" }

func (C *CPU) op_SBC_IM(mem *Memory)  { C.opName = "ToDO" }
func (C *CPU) op_SBC_ZP(mem *Memory)  { C.opName = "ToDO" }
func (C *CPU) op_SBC_ZPX(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABS(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABX(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_ABY(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_INX(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_SBC_INY(mem *Memory) { C.opName = "ToDO" }

func (C *CPU) op_CMP_IM(mem *Memory)  { C.opName = "ToDO" }
func (C *CPU) op_CMP_ZP(mem *Memory)  { C.opName = "ToDO" }
func (C *CPU) op_CMP_ZPX(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_CMP_ABS(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_CMP_ABX(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_CMP_ABY(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_CMP_INX(mem *Memory) { C.opName = "ToDO" }
func (C *CPU) op_CMP_INY(mem *Memory) { C.opName = "ToDO" }

func (C *CPU) op_CPX_IM(mem *Memory) {
	C.opName = "CPX Im"
	value := C.fetchByte(mem)
	C.setC(C.X, value)
	res := C.X - value
	C.setZ(res)
	C.setN(res)
}

func (C *CPU) op_CPX_ZP(mem *Memory)  { C.opName = "ToDO" }
func (C *CPU) op_CPX_ABS(mem *Memory) { C.opName = "ToDO" }

func (C *CPU) op_CPY_IM(mem *Memory)  { C.opName = "ToDO" }
func (C *CPU) op_CPY_ZP(mem *Memory)  { C.opName = "ToDO" }
func (C *CPU) op_CPY_ABS(mem *Memory) { C.opName = "ToDO" }
