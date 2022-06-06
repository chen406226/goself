// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATtiny43U.atdf, see http://packs.download.atmel.com/

//go:build avr && attiny43u
// +build avr,attiny43u

// Device information for the ATtiny43U.
package avr

import (
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATtiny43U"
	ARCH   = "AVR8"
	FAMILY = "tinyAVR"
)

// Interrupts
const (
	IRQ_RESET      = 0  // External Pin, Power-on Reset, Brown-out Reset,Watchdog Reset
	IRQ_INT0       = 1  // External Interrupt Request 0
	IRQ_PCINT0     = 2  // Pin Change Interrupt Request 0
	IRQ_PCINT1     = 3  // Pin Change Interrupt Request 1
	IRQ_WDT        = 4  // Watchdog Time-out
	IRQ_TIM1_COMPA = 5  // Timer/Counter1 Compare Match A
	IRQ_TIM1_COMPB = 6  // Timer/Counter1 Compare Match B
	IRQ_TIM1_OVF   = 7  // Timer/Counter1 Overflow
	IRQ_TIM0_COMPA = 8  // Timer/Counter0 Compare Match A
	IRQ_TIM0_COMPB = 9  // Timer/Counter0 Compare Match B
	IRQ_TIM0_OVF   = 10 // Timer/Counter0 Overflow
	IRQ_ANA_COMP   = 11 // Analog Comparator
	IRQ_ADC        = 12 // ADC Conversion Complete
	IRQ_EE_RDY     = 13 // EEPROM Ready
	IRQ_USI_START  = 14 // USI START
	IRQ_USI_OVF    = 15 // USI Overflow
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

//export __vector_WDT
//go:interrupt
func interruptWDT() {
	callHandlers(IRQ_WDT)
}

//export __vector_TIM1_COMPA
//go:interrupt
func interruptTIM1_COMPA() {
	callHandlers(IRQ_TIM1_COMPA)
}

//export __vector_TIM1_COMPB
//go:interrupt
func interruptTIM1_COMPB() {
	callHandlers(IRQ_TIM1_COMPB)
}

//export __vector_TIM1_OVF
//go:interrupt
func interruptTIM1_OVF() {
	callHandlers(IRQ_TIM1_OVF)
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

//export __vector_TIM0_OVF
//go:interrupt
func interruptTIM0_OVF() {
	callHandlers(IRQ_TIM0_OVF)
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

//export __vector_EE_RDY
//go:interrupt
func interruptEE_RDY() {
	callHandlers(IRQ_EE_RDY)
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
	PORTA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3b)))
	DDRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3a)))
	PINA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x39)))
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x38)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))

	// Universal Serial Interface
	USIBR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x30)))
	USIDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2f)))
	USISR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	USICR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))

	// Watchdog Timer
	WDTCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))

	// Timer/Counter, 8-bit
	TIMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x59)))
	TIFR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x58)))
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x50)))
	TCCR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	TCNT0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x52)))
	OCR0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x56)))
	OCR0B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5c)))
	TIMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))
	TIFR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4f)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))
	TCNT1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	OCR1A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))
	OCR1B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))

	// Bootloader
	SPMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))

	// CPU Registers
	PRR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x20)))
	OSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x51)))
	CLKPR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	SREG   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	SPL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	MCUSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x54)))
	GPIOR2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))
	GPIOR1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x34)))
	GPIOR0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x33)))

	// External Interrupts
	GIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5b)))
	GIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5a)))
	PCMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	PCMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))

	// Analog Comparator
	ACSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))

	// Analog-to-Digital Converter
	ADMUX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))

	// EEPROM
	EEAR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	EEDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EECR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_SELFPRGEN = 0x1 // Self Programming enable

	// HIGH
	HIGH_RSTDISBL  = 0x80 // Reset Disabled (Enable PB7 as i/o pin)
	HIGH_DWEN      = 0x40 // Debug Wire enable
	HIGH_SPIEN     = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON     = 0x10 // Watch-dog Timer always on
	HIGH_EESAVE    = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BODLEVEL0 = 0x1  // Brown-out Detector trigger level
	HIGH_BODLEVEL1 = 0x2  // Brown-out Detector trigger level
	HIGH_BODLEVEL2 = 0x4  // Brown-out Detector trigger level

	// LOW
	LOW_CKDIV8     = 0x80 // Divide clock by 8 internally
	LOW_CKOUT      = 0x40 // Clock output on PORTB5
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

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control Register
	WDTCSR_WDIF = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCSR_WDIE = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCSR_WDP0 = 0x1  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP1 = 0x2  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP2 = 0x4  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP3 = 0x20 // Watchdog Timer Prescaler Bits
	WDTCSR_WDCE = 0x10 // Watchdog Change Enable
	WDTCSR_WDE  = 0x8  // Watch Dog Enable
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TIMSK0: Timer/Counter Interrupt Mask Register
	TIMSK0_OCIE0B = 0x4 // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK0_OCIE0A = 0x2 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK0_TOIE0  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter0 Interrupt Flag Register
	TIFR0_OCF0B = 0x4 // Timer/Counter0 Output Compare Flag B
	TIFR0_OCF0A = 0x2 // Timer/Counter0 Output Compare Flag A
	TIFR0_TOV0  = 0x1 // Timer/Counter0 Overflow Flag

	// TCCR0A: Timer/Counter  Control Register A
	TCCR0A_COM0A0 = 0x40 // Compare Match Output A Mode bits
	TCCR0A_COM0A1 = 0x80 // Compare Match Output A Mode bits
	TCCR0A_COM0B0 = 0x10 // Compare Match Output B Mode bits
	TCCR0A_COM0B1 = 0x20 // Compare Match Output B Mode bits
	TCCR0A_WGM00  = 0x1  // Waveform Generation Mode bits
	TCCR0A_WGM01  = 0x2  // Waveform Generation Mode bits

	// TCCR0B: Timer/Counter Control Register B
	TCCR0B_FOC0A = 0x80 // Force Output Compare A
	TCCR0B_FOC0B = 0x40 // Force Output Compare B
	TCCR0B_WGM02 = 0x8  // Waveform Generation Mode bit 2
	TCCR0B_CS00  = 0x1  // Clock Select bits
	TCCR0B_CS01  = 0x2  // Clock Select bits
	TCCR0B_CS02  = 0x4  // Clock Select bits

	// TIMSK1: Timer/Counter Interrupt Mask Register
	TIMSK1_OCIE1B = 0x4 // Timer/Counter1 Output Compare Match B Interrupt Enable
	TIMSK1_OCIE1A = 0x2 // Timer/Counter1 Output Compare Match A Interrupt Enable
	TIMSK1_TOIE1  = 0x1 // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter1 Interrupt Flag Register
	TIFR1_OCF1B = 0x4 // Timer/Counter1 Output Compare Flag B
	TIFR1_OCF1A = 0x2 // Timer/Counter1 Output Compare Flag A
	TIFR1_TOV1  = 0x1 // Timer/Counter1 Overflow Flag

	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A0 = 0x40 // Compare Match Output A Mode bits
	TCCR1A_COM1A1 = 0x80 // Compare Match Output A Mode bits
	TCCR1A_COM1B0 = 0x10 // Compare Match Output B Mode bits
	TCCR1A_COM1B1 = 0x20 // Compare Match Output B Mode bits
	TCCR1A_WGM10  = 0x1  // Waveform Generation Mode bits
	TCCR1A_WGM11  = 0x2  // Waveform Generation Mode bits

	// TCCR1B: Timer/Counter Control Register B
	TCCR1B_FOC1A = 0x80 // Force Output Compare A
	TCCR1B_FOC1B = 0x40 // Force Output Compare B
	TCCR1B_WGM12 = 0x8  // Waveform Generation Mode bit 2
	TCCR1B_CS10  = 0x1  // Clock Select bits
	TCCR1B_CS11  = 0x2  // Clock Select bits
	TCCR1B_CS12  = 0x4  // Clock Select bits
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_CTPB  = 0x10 // Clear temporary page buffer
	SPMCSR_RFLB  = 0x8  // Read fuse and lock bits
	SPMCSR_PGWRT = 0x4  // Page Write
	SPMCSR_PGERS = 0x2  // Page Erase
	SPMCSR_SPMEN = 0x1  // Store Program Memory Enable
)

// Bitfields for CPU: CPU Registers
const (
	// PRR: Power Reduction Register
	PRR_PRTIM1 = 0x8 // Power Reduction Timer/Counter1
	PRR_PRTIM0 = 0x4 // Power Reduction Timer/Counter0
	PRR_PRUSI  = 0x2 // Power Reduction USI
	PRR_PRADC  = 0x1 // Power Reduction ADC

	// OSCCAL: Oscillator Calibration Value
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

	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// MCUSR: MCU Status Register
	MCUSR_WDRF  = 0x8 // Watchdog Reset Flag
	MCUSR_BORF  = 0x4 // Brown-out Reset Flag
	MCUSR_EXTRF = 0x2 // External Reset Flag
	MCUSR_PORF  = 0x1 // Power-on reset flag
)

// Bitfields for EXINT: External Interrupts
const (
	// GIMSK: General Interrupt Mask Register
	GIMSK_INT0  = 0x40 // External Interrupt Request 0 Enable
	GIMSK_PCIE0 = 0x10 // Pin Change Interrupt Enables
	GIMSK_PCIE1 = 0x20 // Pin Change Interrupt Enables

	// GIFR: General Interrupt Flag register
	GIFR_INTF0 = 0x40 // External Interrupt Flag 0
	GIFR_PCIF0 = 0x10 // Pin Change Interrupt Flags
	GIFR_PCIF1 = 0x20 // Pin Change Interrupt Flags
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
	// ADMUX: ADC Multiplexer Selection Register
	ADMUX_REFS = 0x40 // Reference Selection Bit
	ADMUX_MUX0 = 0x1  // Analog Channel Selection Bits
	ADMUX_MUX1 = 0x2  // Analog Channel Selection Bits
	ADMUX_MUX2 = 0x4  // Analog Channel Selection Bits

	// ADCSRA: ADC Control and Status Register A
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE = 0x20 // ADC Auto Trigger Enable
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS0 = 0x1  // ADC  Prescaler Select Bits
	ADCSRA_ADPS1 = 0x2  // ADC  Prescaler Select Bits
	ADCSRA_ADPS2 = 0x4  // ADC  Prescaler Select Bits
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