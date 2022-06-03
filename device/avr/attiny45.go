// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATtiny45.atdf, see http://packs.download.atmel.com/

//go:build avr && attiny45
// +build avr,attiny45

// Device information for the ATtiny45.
package avr

import (
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATtiny45"
	ARCH   = "AVR8"
	FAMILY = "tinyAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin, Power-on Reset, Brown-out Reset,Watchdog Reset
	IRQ_INT0         = 1  // External Interrupt 0
	IRQ_PCINT0       = 2  // Pin change Interrupt Request 0
	IRQ_TIMER1_COMPA = 3  // Timer/Counter1 Compare Match 1A
	IRQ_TIMER1_OVF   = 4  // Timer/Counter1 Overflow
	IRQ_TIMER0_OVF   = 5  // Timer/Counter0 Overflow
	IRQ_EE_RDY       = 6  // EEPROM Ready
	IRQ_ANA_COMP     = 7  // Analog comparator
	IRQ_ADC          = 8  // ADC Conversion ready
	IRQ_TIMER1_COMPB = 9  // Timer/Counter1 Compare Match B
	IRQ_TIMER0_COMPA = 10 // Timer/Counter0 Compare Match A
	IRQ_TIMER0_COMPB = 11 // Timer/Counter0 Compare Match B
	IRQ_WDT          = 12 // Watchdog Time-out
	IRQ_USI_START    = 13 // USI START
	IRQ_USI_OVF      = 14 // USI Overflow
	IRQ_max          = 14 // Highest interrupt number on this device.
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

//export __vector_TIMER1_COMPA
//go:interrupt
func interruptTIMER1_COMPA() {
	callHandlers(IRQ_TIMER1_COMPA)
}

//export __vector_TIMER1_OVF
//go:interrupt
func interruptTIMER1_OVF() {
	callHandlers(IRQ_TIMER1_OVF)
}

//export __vector_TIMER0_OVF
//go:interrupt
func interruptTIMER0_OVF() {
	callHandlers(IRQ_TIMER0_OVF)
}

//export __vector_EE_RDY
//go:interrupt
func interruptEE_RDY() {
	callHandlers(IRQ_EE_RDY)
}

//export __vector_ANA_COMP
//go:interrupt
func interruptANA_COMP() {
	callHandlers(IRQ_ANA_COMP)
}

//export __vector_ADC
//go:interrupt
func interruptADC() {
	callHandlers(IRQ_ADC)
}

//export __vector_TIMER1_COMPB
//go:interrupt
func interruptTIMER1_COMPB() {
	callHandlers(IRQ_TIMER1_COMPB)
}

//export __vector_TIMER0_COMPA
//go:interrupt
func interruptTIMER0_COMPA() {
	callHandlers(IRQ_TIMER0_COMPA)
}

//export __vector_TIMER0_COMPB
//go:interrupt
func interruptTIMER0_COMPB() {
	callHandlers(IRQ_TIMER0_COMPB)
}

//export __vector_WDT
//go:interrupt
func interruptWDT() {
	callHandlers(IRQ_WDT)
}

//export __vector_USI_START
//go:interrupt
func interruptUSI_START() {
	callHandlers(IRQ_USI_START)
}

//export __vector_USI_OVF
//go:interrupt
func interruptUSI_OVF() {
	callHandlers(IRQ_USI_OVF)
}

// Peripherals.
var (
	// Fuses
	EXTENDED = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2)))
	HIGH     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1)))
	LOW      = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Lockbits
	LOCKBIT = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// I/O Port
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x38)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))

	// Analog Comparator
	ACSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))

	// Analog-to-Digital Converter
	ADMUX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))

	// Universal Serial Interface
	USIBR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x30)))
	USIDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2f)))
	USISR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	USICR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))

	// External Interrupts
	GIMSK = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5b)))
	GIFR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5a)))
	PCMSK = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	EEARH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))

	// Watchdog Timer
	WDTCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))

	// Timer/Counter, 8-bit
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	TCCR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	TCNT0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x52)))
	OCR0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x49)))
	OCR0B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x48)))
	TCCR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x50)))
	TCNT1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4f)))
	OCR1A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))
	OCR1B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	OCR1C  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	DTPS   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x43)))
	DT1A   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x45)))
	DT1B   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x44)))

	// CPU Registers
	SREG   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	PRR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	SPL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	MCUSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x54)))
	OSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x51)))
	CLKPR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	PLLCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x47)))
	DWDR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x42)))
	GPIOR2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x33)))
	GPIOR1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))
	GPIOR0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x31)))

	// Bootloader
	SPMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_SELFPRGEN = 0x1 // Self Programming enable

	// HIGH
	HIGH_RSTDISBL  = 0x80 // Reset Disabled (Enable PB5 as i/o pin)
	HIGH_DWEN      = 0x40 // Debug Wire enable
	HIGH_SPIEN     = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON     = 0x10 // Watch-dog Timer always on
	HIGH_EESAVE    = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BODLEVEL0 = 0x1  // Brown-out Detector trigger level
	HIGH_BODLEVEL1 = 0x2  // Brown-out Detector trigger level
	HIGH_BODLEVEL2 = 0x4  // Brown-out Detector trigger level

	// LOW
	LOW_CKDIV8     = 0x80 // Divide clock by 8 internally
	LOW_CKOUT      = 0x40 // Clock output on PORTB4
	LOW_SUT_CKSEL0 = 0x1  // Select Clock source
	LOW_SUT_CKSEL1 = 0x2  // Select Clock source
	LOW_SUT_CKSEL2 = 0x4  // Select Clock source
	LOW_SUT_CKSEL3 = 0x8  // Select Clock source
	LOW_SUT_CKSEL4 = 0x10 // Select Clock source
	LOW_SUT_CKSEL5 = 0x20 // Select Clock source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB0 = 0x1 // Memory Lock
	LOCKBIT_LB1 = 0x2 // Memory Lock
)

// Bitfields for AC: Analog Comparator
const (
	// ACSR: Analog Comparator Control And Status Register
	ACSR_ACD   = 0x80 // Analog Comparator Disable
	ACSR_ACBG  = 0x40 // Analog Comparator Bandgap Select
	ACSR_ACO   = 0x20 // Analog Compare Output
	ACSR_ACI   = 0x10 // Analog Comparator Interrupt Flag
	ACSR_ACIE  = 0x8  // Analog Comparator Interrupt Enable
	ACSR_ACIS0 = 0x1  // Analog Comparator Interrupt Mode Select bits
	ACSR_ACIS1 = 0x2  // Analog Comparator Interrupt Mode Select bits
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS0 = 0x40 // Reference Selection Bits
	ADMUX_REFS1 = 0x80 // Reference Selection Bits
	ADMUX_ADLAR = 0x20 // Left Adjust Result
	ADMUX_REFS2 = 0x10 // Reference Selection Bit 2
	ADMUX_MUX0  = 0x1  // Analog Channel and Gain Selection Bits
	ADMUX_MUX1  = 0x2  // Analog Channel and Gain Selection Bits
	ADMUX_MUX2  = 0x4  // Analog Channel and Gain Selection Bits
	ADMUX_MUX3  = 0x8  // Analog Channel and Gain Selection Bits

	// ADCSRA: The ADC Control and Status register
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE = 0x20 // ADC Auto Trigger Enable
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS0 = 0x1  // ADC  Prescaler Select Bits
	ADCSRA_ADPS1 = 0x2  // ADC  Prescaler Select Bits
	ADCSRA_ADPS2 = 0x4  // ADC  Prescaler Select Bits
)

// Bitfields for USI: Universal Serial Interface
const (
	// USISR: USI Status Register
	USISR_USISIF  = 0x80 // Start Condition Interrupt Flag
	USISR_USIOIF  = 0x40 // Counter Overflow Interrupt Flag
	USISR_USIPF   = 0x20 // Stop Condition Flag
	USISR_USIDC   = 0x10 // Data Output Collision
	USISR_USICNT0 = 0x1  // USI Counter Value Bits
	USISR_USICNT1 = 0x2  // USI Counter Value Bits
	USISR_USICNT2 = 0x4  // USI Counter Value Bits
	USISR_USICNT3 = 0x8  // USI Counter Value Bits

	// USICR: USI Control Register
	USICR_USISIE = 0x80 // Start Condition Interrupt Enable
	USICR_USIOIE = 0x40 // Counter Overflow Interrupt Enable
	USICR_USIWM0 = 0x10 // USI Wire Mode Bits
	USICR_USIWM1 = 0x20 // USI Wire Mode Bits
	USICR_USICS0 = 0x4  // USI Clock Source Select Bits
	USICR_USICS1 = 0x8  // USI Clock Source Select Bits
	USICR_USICLK = 0x2  // Clock Strobe
	USICR_USITC  = 0x1  // Toggle Clock Port Pin
)

// Bitfields for EXINT: External Interrupts
const (
	// GIMSK: General Interrupt Mask Register
	GIMSK_INT0 = 0x40 // External Interrupt Request 0 Enable
	GIMSK_PCIE = 0x20 // Pin Change Interrupt Enable

	// GIFR: General Interrupt Flag register
	GIFR_INTF0 = 0x40 // External Interrupt Flag 0
	GIFR_PCIF  = 0x20 // Pin Change Interrupt Flag
)

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EEPM0 = 0x10 // EEPROM Programming Mode Bits
	EECR_EEPM1 = 0x20 // EEPROM Programming Mode Bits
	EECR_EERIE = 0x8  // EEPROM Ready Interrupt Enable
	EECR_EEMPE = 0x4  // EEPROM Master Write Enable
	EECR_EEPE  = 0x2  // EEPROM Write Enable
	EECR_EERE  = 0x1  // EEPROM Read Enable
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCR: Watchdog Timer Control Register
	WDTCR_WDIF = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCR_WDIE = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCR_WDP0 = 0x1  // Watchdog Timer Prescaler Bits
	WDTCR_WDP1 = 0x2  // Watchdog Timer Prescaler Bits
	WDTCR_WDP2 = 0x4  // Watchdog Timer Prescaler Bits
	WDTCR_WDP3 = 0x20 // Watchdog Timer Prescaler Bits
	WDTCR_WDCE = 0x10 // Watchdog Change Enable
	WDTCR_WDE  = 0x8  // Watch Dog Enable
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR0A: Timer/Counter  Control Register A
	TCCR0A_COM0A0 = 0x40 // Compare Output Mode, Phase Correct PWM Mode
	TCCR0A_COM0A1 = 0x80 // Compare Output Mode, Phase Correct PWM Mode
	TCCR0A_COM0B0 = 0x10 // Compare Output Mode, Fast PWm
	TCCR0A_COM0B1 = 0x20 // Compare Output Mode, Fast PWm
	TCCR0A_WGM00  = 0x1  // Waveform Generation Mode
	TCCR0A_WGM01  = 0x2  // Waveform Generation Mode

	// TCCR0B: Timer/Counter Control Register B
	TCCR0B_FOC0A = 0x80 // Force Output Compare A
	TCCR0B_FOC0B = 0x40 // Force Output Compare B
	TCCR0B_WGM02 = 0x8
	TCCR0B_CS00  = 0x1 // Clock Select
	TCCR0B_CS01  = 0x2 // Clock Select
	TCCR0B_CS02  = 0x4 // Clock Select

	// TCCR1: Timer/Counter Control Register
	TCCR1_CTC1   = 0x80 // Clear Timer/Counter on Compare Match
	TCCR1_PWM1A  = 0x40 // Pulse Width Modulator Enable
	TCCR1_COM1A0 = 0x10 // Compare Output Mode, Bits
	TCCR1_COM1A1 = 0x20 // Compare Output Mode, Bits
	TCCR1_CS10   = 0x1  // Clock Select Bits
	TCCR1_CS11   = 0x2  // Clock Select Bits
	TCCR1_CS12   = 0x4  // Clock Select Bits
	TCCR1_CS13   = 0x8  // Clock Select Bits

	// DTPS: Dead time prescaler register
	DTPS_DTPS0 = 0x1
	DTPS_DTPS1 = 0x2

	// DT1A: Dead time value register
	DT1A_DTVH0 = 0x10
	DT1A_DTVH1 = 0x20
	DT1A_DTVH2 = 0x40
	DT1A_DTVH3 = 0x80
	DT1A_DTVL0 = 0x1
	DT1A_DTVL1 = 0x2
	DT1A_DTVL2 = 0x4
	DT1A_DTVL3 = 0x8

	// DT1B: Dead time value B
	DT1B_DTVH0 = 0x10
	DT1B_DTVH1 = 0x20
	DT1B_DTVH2 = 0x40
	DT1B_DTVH3 = 0x80
	DT1B_DTVL0 = 0x1
	DT1B_DTVL1 = 0x2
	DT1B_DTVL2 = 0x4
	DT1B_DTVL3 = 0x8
)

// Bitfields for CPU: CPU Registers
const (
	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// PRR: Power Reduction Register
	PRR_PRTIM1 = 0x8 // Power Reduction Timer/Counter1
	PRR_PRTIM0 = 0x4 // Power Reduction Timer/Counter0
	PRR_PRUSI  = 0x2 // Power Reduction USI
	PRR_PRADC  = 0x1 // Power Reduction ADC

	// MCUSR: MCU Status register
	MCUSR_WDRF  = 0x8 // Watchdog Reset Flag
	MCUSR_BORF  = 0x4 // Brown-out Reset Flag
	MCUSR_EXTRF = 0x2 // External Reset Flag
	MCUSR_PORF  = 0x1 // Power-On Reset Flag

	// OSCCAL: Oscillator Calibration Register
	OSCCAL_OSCCAL0 = 0x1  // Oscillator Calibration
	OSCCAL_OSCCAL1 = 0x2  // Oscillator Calibration
	OSCCAL_OSCCAL2 = 0x4  // Oscillator Calibration
	OSCCAL_OSCCAL3 = 0x8  // Oscillator Calibration
	OSCCAL_OSCCAL4 = 0x10 // Oscillator Calibration
	OSCCAL_OSCCAL5 = 0x20 // Oscillator Calibration
	OSCCAL_OSCCAL6 = 0x40 // Oscillator Calibration
	OSCCAL_OSCCAL7 = 0x80 // Oscillator Calibration

	// CLKPR: Clock Prescale Register
	CLKPR_CLKPCE = 0x80 // Clock Prescaler Change Enable
	CLKPR_CLKPS0 = 0x1  // Clock Prescaler Select Bits
	CLKPR_CLKPS1 = 0x2  // Clock Prescaler Select Bits
	CLKPR_CLKPS2 = 0x4  // Clock Prescaler Select Bits
	CLKPR_CLKPS3 = 0x8  // Clock Prescaler Select Bits

	// PLLCSR: PLL Control and status register
	PLLCSR_LSM   = 0x80 // Low speed mode
	PLLCSR_PCKE  = 0x4  // PCK Enable
	PLLCSR_PLLE  = 0x2  // PLL Enable
	PLLCSR_PLOCK = 0x1  // PLL Lock detector
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_RSIG  = 0x20 // Read Device Signature Imprint Table
	SPMCSR_CTPB  = 0x10 // Clear temporary page buffer
	SPMCSR_RFLB  = 0x8  // Read fuse and lock bits
	SPMCSR_PGWRT = 0x4  // Page Write
	SPMCSR_PGERS = 0x2  // Page Erase
	SPMCSR_SPMEN = 0x1  // Store Program Memory Enable
)
