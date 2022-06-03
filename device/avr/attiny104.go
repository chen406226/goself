// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATtiny104.atdf, see http://packs.download.atmel.com/

//go:build avr && attiny104
// +build avr,attiny104

// Device information for the ATtiny104.
package avr

import (
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATtiny104"
	ARCH   = "AVR8L"
	FAMILY = "tinyAVR"
)

// Interrupts
const (
	IRQ_RESET      = 0  // External Reset, Power-on Reset and Watchdog Reset
	IRQ_INT0       = 1  // External Interrupt Request 0
	IRQ_PCINT0     = 2  // Pin Change Interrupt Request 0
	IRQ_PCINT1     = 3  // Pin Change Interrupt Request 1
	IRQ_TIM0_CAPT  = 4  // Timer/Counter0 Input Capture
	IRQ_TIM0_OVF   = 5  // Timer/Counter0 Overflow
	IRQ_TIM0_COMPA = 6  // Timer/Counter Compare Match A
	IRQ_TIM0_COMPB = 7  // Timer/Counter Compare Match B
	IRQ_ANA_COMP   = 8  // Analog Comparator
	IRQ_WDT        = 9  // Watchdog Time-out
	IRQ_VLM        = 10 // Vcc Voltage Level Monitor
	IRQ_ADC        = 11 // ADC Conversion complete
	IRQ_USART_RXS  = 12 // USART RX Start
	IRQ_USART_RXC  = 13 // USART RX Complete
	IRQ_USART_DRE  = 14 // USART Data register empty
	IRQ_USART_TXC  = 15 // USART Tx Complete
	IRQ_max        = 15 // Highest interrupt number on this device.
)

// Pseudo function call that is replaced by the compiler with the actual
// functions registered through interrupt.New.
//go:linkname callHandlers runtime/interrupt.callHandlers
func callHandlers(num int)

//export __vector_RESET
//go:interrupt
func interruptRESET() {
	callHandlers(IRQ_RESET)
}

//export __vector_INT0
//go:interrupt
func interruptINT0() {
	callHandlers(IRQ_INT0)
}

//export __vector_PCINT0
//go:interrupt
func interruptPCINT0() {
	callHandlers(IRQ_PCINT0)
}

//export __vector_PCINT1
//go:interrupt
func interruptPCINT1() {
	callHandlers(IRQ_PCINT1)
}

//export __vector_TIM0_CAPT
//go:interrupt
func interruptTIM0_CAPT() {
	callHandlers(IRQ_TIM0_CAPT)
}

//export __vector_TIM0_OVF
//go:interrupt
func interruptTIM0_OVF() {
	callHandlers(IRQ_TIM0_OVF)
}

//export __vector_TIM0_COMPA
//go:interrupt
func interruptTIM0_COMPA() {
	callHandlers(IRQ_TIM0_COMPA)
}

//export __vector_TIM0_COMPB
//go:interrupt
func interruptTIM0_COMPB() {
	callHandlers(IRQ_TIM0_COMPB)
}

//export __vector_ANA_COMP
//go:interrupt
func interruptANA_COMP() {
	callHandlers(IRQ_ANA_COMP)
}

//export __vector_WDT
//go:interrupt
func interruptWDT() {
	callHandlers(IRQ_WDT)
}

//export __vector_VLM
//go:interrupt
func interruptVLM() {
	callHandlers(IRQ_VLM)
}

//export __vector_ADC
//go:interrupt
func interruptADC() {
	callHandlers(IRQ_ADC)
}

//export __vector_USART_RXS
//go:interrupt
func interruptUSART_RXS() {
	callHandlers(IRQ_USART_RXS)
}

//export __vector_USART_RXC
//go:interrupt
func interruptUSART_RXC() {
	callHandlers(IRQ_USART_RXC)
}

//export __vector_USART_DRE
//go:interrupt
func interruptUSART_DRE() {
	callHandlers(IRQ_USART_DRE)
}

//export __vector_USART_TXC
//go:interrupt
func interruptUSART_TXC() {
	callHandlers(IRQ_USART_TXC)
}

// Peripherals.
var (
	// Fuses
	BYTE0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Lockbits
	LOCKBIT = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Analog-to-Digital Converter
	ADMUX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1b)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x19)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1a)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1d)))
	ADCSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1c)))

	// Analog Comparator
	ACSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1f)))
	ACSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1e)))

	// CPU Registers
	CCP    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))
	SPL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	SPH    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	SREG   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))
	CLKMSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	CLKPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))
	OSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x39)))
	SMCR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3a)))
	PRR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))
	VLMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x34)))
	RSTFLR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3b)))
	NVMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))
	NVMCMD = (*volatile.Register8)(unsafe.Pointer(uintptr(0x33)))

	// I/O Port
	PUEB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4)))
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6)))
	PUEA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3)))
	DDRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1)))
	PINA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))
	PORTA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2)))

	// External Interrupts
	EICRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x15)))
	EIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x13)))
	EIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x14)))
	PCICR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x12)))
	PCIFR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x11)))
	PCMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x10)))
	PCMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf)))

	// Timer/Counter, 16-bit
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	TCCR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))
	TCCR0C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))
	TCNT0L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))
	TCNT0H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x29)))
	OCR0AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	OCR0AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	OCR0BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	OCR0BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
	ICR0L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x22)))
	ICR0H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	TIMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	TIFR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2a)))
	GTCCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2f)))

	// USART
	UDR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8)))
	UCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe)))
	UCSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0xd)))
	UCSRC = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc)))
	UCSRD = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb)))
	UBRRL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x9)))
	UBRRH = (*volatile.Register8)(unsafe.Pointer(uintptr(0xa)))

	// Watchdog Timer
	WDTCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x31)))
)

// Bitfields for FUSE: Fuses
const (
	// BYTE0
	BYTE0_SELFPROGEN = 0x8 // Self programming enable
	BYTE0_CKOUT      = 0x4 // Output external clock
	BYTE0_WDTON      = 0x2 // Watch dog timer always on
	BYTE0_RSTDISBL   = 0x1 // Disable external reset
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB0 = 0x1 // Memory Lock
	LOCKBIT_LB1 = 0x2 // Memory Lock
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_MUX0  = 0x1  // Analog Channel Selection Bits
	ADMUX_MUX1  = 0x2  // Analog Channel Selection Bits
	ADMUX_MUX2  = 0x4  // Analog Channel Selection Bits
	ADMUX_REFS0 = 0x40 // Analog Reference voltage Selection Bits
	ADMUX_REFS1 = 0x80 // Analog Reference voltage Selection Bits

	// ADCSRA: The ADC Control and Status register A
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE = 0x20 // ADC  Auto Trigger Enable
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS0 = 0x1  // ADC  Prescaler Select Bits
	ADCSRA_ADPS1 = 0x2  // ADC  Prescaler Select Bits
	ADCSRA_ADPS2 = 0x4  // ADC  Prescaler Select Bits

	// ADCSRB: The ADC Control and Status register B
	ADCSRB_ADLAR = 0x80 // Left Adjustment for ADC Result Readout
	ADCSRB_ADTS0 = 0x1  // ADC Auto Trigger Source bits
	ADCSRB_ADTS1 = 0x2  // ADC Auto Trigger Source bits
	ADCSRB_ADTS2 = 0x4  // ADC Auto Trigger Source bits
)

// Bitfields for AC: Analog Comparator
const (
	// ACSRA: Analog Comparator Control And Status Register A
	ACSRA_ACD   = 0x80 // Analog Comparator Disable
	ACSRA_ACBG  = 0x40 // Analog Comparator Bandgap enable
	ACSRA_ACO   = 0x20 // Analog Compare Output
	ACSRA_ACI   = 0x10 // Analog Comparator Interrupt Flag
	ACSRA_ACIE  = 0x8  // Analog Comparator Interrupt Enable
	ACSRA_ACIC  = 0x4  // Analog Comparator Input Capture  Enable
	ACSRA_ACIS0 = 0x1  // Analog Comparator Interrupt Mode Select bits
	ACSRA_ACIS1 = 0x2  // Analog Comparator Interrupt Mode Select bits

	// ACSRB: Analog Comparator Control And Status Register B
	ACSRB_ACOE   = 0x2 // Analog Comparator Output Enable
	ACSRB_ACPMUX = 0x1 // Analog Comparator positive input selection bit
)

// Bitfields for CPU: CPU Registers
const (
	// CCP: Configuration Change Protection
	CCP_CCP0 = 0x1  // CCP signature
	CCP_CCP1 = 0x2  // CCP signature
	CCP_CCP2 = 0x4  // CCP signature
	CCP_CCP3 = 0x8  // CCP signature
	CCP_CCP4 = 0x10 // CCP signature
	CCP_CCP5 = 0x20 // CCP signature
	CCP_CCP6 = 0x40 // CCP signature
	CCP_CCP7 = 0x80 // CCP signature

	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// CLKMSR: Clock Main Settings Register
	CLKMSR_CLKMS0 = 0x1 // Clock Main Select Bits
	CLKMSR_CLKMS1 = 0x2 // Clock Main Select Bits

	// CLKPSR: Clock Prescale Register
	CLKPSR_CLKPS0 = 0x1 // Clock Prescaler Select Bits
	CLKPSR_CLKPS1 = 0x2 // Clock Prescaler Select Bits
	CLKPSR_CLKPS2 = 0x4 // Clock Prescaler Select Bits
	CLKPSR_CLKPS3 = 0x8 // Clock Prescaler Select Bits

	// SMCR: Sleep Mode Control Register
	SMCR_SM0 = 0x2 // Sleep Mode Select Bits
	SMCR_SM1 = 0x4 // Sleep Mode Select Bits
	SMCR_SM2 = 0x8 // Sleep Mode Select Bits
	SMCR_SE  = 0x1 // Sleep Enable

	// PRR: Power Reduction Register
	PRR_PRUSART = 0x4 // Power Reduction USART
	PRR_PRADC   = 0x2 // Power Reduction ADC
	PRR_PRTIM0  = 0x1 // Power Reduction Timer/Counter0

	// VLMCSR: Vcc Level Monitoring Control and Status Register
	VLMCSR_VLMF  = 0x80 // VLM Flag
	VLMCSR_VLMIE = 0x40 // VLM Interrupt Enable
	VLMCSR_VLM0  = 0x1  // Trigger Level of Voltage Level Monitor bits
	VLMCSR_VLM1  = 0x2  // Trigger Level of Voltage Level Monitor bits
	VLMCSR_VLM2  = 0x4  // Trigger Level of Voltage Level Monitor bits

	// RSTFLR: Reset Flag Register
	RSTFLR_WDRF  = 0x8 // Watchdog Reset Flag
	RSTFLR_EXTRF = 0x2 // External Reset Flag
	RSTFLR_PORF  = 0x1 // Power-on Reset Flag

	// NVMCSR: Non-Volatile Memory Control and Status Register
	NVMCSR_NVMBSY = 0x80 // Non-Volatile Memory Busy
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register A
	EICRA_ISC01 = 0x2 // Interrupt Sense Control 0 Bit 1
	EICRA_ISC00 = 0x1 // Interrupt Sense Control 0 Bit 0

	// EIMSK: External Interrupt Mask register
	EIMSK_INT0 = 0x1 // External Interrupt Request 0 Enable

	// EIFR: External Interrupt Flag register
	EIFR_INTF0 = 0x1 // External Interrupt Flag 0

	// PCICR: Pin Change Interrupt Control Register
	PCICR_PCIE1 = 0x2 // Pin Change Interrupt Enable 1
	PCICR_PCIE0 = 0x1 // Pin Change Interrupt Enable 0

	// PCIFR: Pin Change Interrupt Flag Register
	PCIFR_PCIF1 = 0x2 // Pin Change Interrupt Flag 1
	PCIFR_PCIF0 = 0x1 // Pin Change Interrupt Flag 0

	// PCMSK1: Pin Change Mask Register 1
	PCMSK1_PCINT11 = 0x8 // Pin Change Enable Mask 1 Bit 3
	PCMSK1_PCINT10 = 0x4 // Pin Change Enable Mask 1 Bit 2
	PCMSK1_PCINT9  = 0x2 // Pin Change Enable Mask 1 Bit 1
	PCMSK1_PCINT8  = 0x1 // Pin Change Enable Mask 1 Bit 0

	// PCMSK0: Pin Change Mask Register 0
	PCMSK0_PCINT7 = 0x80 // Pin Change Enable Mask 0 Bit 7
	PCMSK0_PCINT6 = 0x40 // Pin Change Enable Mask 0 Bit 6
	PCMSK0_PCINT5 = 0x20 // Pin Change Enable Mask 0 Bit 5
	PCMSK0_PCINT4 = 0x10 // Pin Change Enable Mask 0 Bit 4
	PCMSK0_PCINT3 = 0x8  // Pin Change Enable Mask 0 Bit 3
	PCMSK0_PCINT2 = 0x4  // Pin Change Enable Mask 0 Bit 2
	PCMSK0_PCINT1 = 0x2  // Pin Change Enable Mask 0 Bit 1
	PCMSK0_PCINT0 = 0x1  // Pin Change Enable Mask 0 Bit 0
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR0A: Timer/Counter 0 Control Register A
	TCCR0A_COM0A0 = 0x40 // Compare Output Mode for Channel A bits
	TCCR0A_COM0A1 = 0x80 // Compare Output Mode for Channel A bits
	TCCR0A_COM0B0 = 0x10 // Compare Output Mode for Channel B bits
	TCCR0A_COM0B1 = 0x20 // Compare Output Mode for Channel B bits
	TCCR0A_WGM00  = 0x1  // Waveform Generation Mode
	TCCR0A_WGM01  = 0x2  // Waveform Generation Mode

	// TCCR0B: Timer/Counter 0 Control Register B
	TCCR0B_ICNC0 = 0x80 // Input Capture Noise Canceler
	TCCR0B_ICES0 = 0x40 // Input Capture Edge Select
	TCCR0B_WGM00 = 0x8  // Waveform Generation Mode
	TCCR0B_WGM01 = 0x10 // Waveform Generation Mode
	TCCR0B_CS00  = 0x1  // Clock Select
	TCCR0B_CS01  = 0x2  // Clock Select
	TCCR0B_CS02  = 0x4  // Clock Select

	// TCCR0C: Timer/Counter 0 Control Register C
	TCCR0C_FOC0A = 0x80 // Force Output Compare for Channel A
	TCCR0C_FOC0B = 0x40 // Force Output Compare for Channel B

	// TIMSK0: Timer Interrupt Mask Register 0
	TIMSK0_ICIE0  = 0x20 // Input Capture Interrupt Enable
	TIMSK0_OCIE0B = 0x4  // Output Compare B Match Interrupt Enable
	TIMSK0_OCIE0A = 0x2  // Output Compare A Match Interrupt Enable
	TIMSK0_TOIE0  = 0x1  // Overflow Interrupt Enable

	// TIFR0: Overflow Interrupt Enable
	TIFR0_ICF0  = 0x20 // Input Capture Flag
	TIFR0_OCF0B = 0x4  // Timer Output Compare Flag 0B
	TIFR0_OCF0A = 0x2  // Timer Output Compare Flag 0A
	TIFR0_TOV0  = 0x1  // Timer Overflow Flag

	// GTCCR: General Timer/Counter Control Register
	GTCCR_TSM   = 0x80 // Timer Synchronization Mode
	GTCCR_REMAP = 0x2  // Remap Bit for 14 pin part only
	GTCCR_PSR   = 0x1  // Prescaler Reset
)

// Bitfields for USART: USART
const (
	// UCSRA: USART Control and Status Register A
	UCSRA_RXC  = 0x80 // USART Receive Complete
	UCSRA_TXC  = 0x40 // USART Transmitt Complete
	UCSRA_UDRE = 0x20 // USART Data Register Empty
	UCSRA_FE   = 0x10 // Framing Error
	UCSRA_DOR  = 0x8  // Data overRun
	UCSRA_UPE  = 0x4  // Parity Error
	UCSRA_U2X  = 0x2  // Double the USART transmission speed
	UCSRA_MPCM = 0x1  // Multi-processor Communication Mode

	// UCSRB: USART Control and Status Register B
	UCSRB_RXCIE = 0x80 // RX Complete Interrupt Enable
	UCSRB_TXCIE = 0x40 // TX Complete Interrupt Enable
	UCSRB_UDRIE = 0x20 // USART Data register Empty Interrupt Enable
	UCSRB_RXEN  = 0x10 // Receiver Enable
	UCSRB_TXEN  = 0x8  // Transmitter Enable
	UCSRB_UCSZ2 = 0x4  // Character Size
	UCSRB_RXB8  = 0x2  // Receive Data Bit 8
	UCSRB_TXB8  = 0x1  // Transmit Data Bit 8

	// UCSRC: USART Control and Status Register C
	UCSRC_UMSEL0 = 0x40 // USART Mode Select
	UCSRC_UMSEL1 = 0x80 // USART Mode Select
	UCSRC_UPM0   = 0x10 // Parity Mode Bits
	UCSRC_UPM1   = 0x20 // Parity Mode Bits
	UCSRC_USBS   = 0x8  // Stop Bit Select
	UCSRC_UCSZ0  = 0x2  // Character Size
	UCSRC_UCSZ1  = 0x4  // Character Size
	UCSRC_UCPOL  = 0x1  // Clock Polarity

	// UCSRD: USART Control and Status Register D
	UCSRD_RXSIE = 0x80 // USART RX Start Interrupt Enable
	UCSRD_RXS   = 0x40 // USART RX Start Flag
	UCSRD_SFDE  = 0x20 // Start frame detection enable
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control and Status Register
	WDTCSR_WDIF = 0x80 // Watchdog Timer Interrupt Flag
	WDTCSR_WDIE = 0x40 // Watchdog Timer Interrupt Enable
	WDTCSR_WDP0 = 0x1  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP1 = 0x2  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP2 = 0x4  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP3 = 0x20 // Watchdog Timer Prescaler Bits
	WDTCSR_WDE  = 0x8  // Watch Dog Enable
)
