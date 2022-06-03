// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATmega32HVBrevB.atdf, see http://packs.download.atmel.com/

//go:build avr && atmega32hvbrevb
// +build avr,atmega32hvbrevb

// Device information for the ATmega32HVBrevB.
package avr

import (
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATmega32HVBrevB"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET         = 0  // External Pin, Power-on Reset, Brown-out Reset and Watchdog Reset
	IRQ_BPINT         = 1  // Battery Protection Interrupt
	IRQ_VREGMON       = 2  // Voltage regulator monitor interrupt
	IRQ_INT0          = 3  // External Interrupt Request 0
	IRQ_INT1          = 4  // External Interrupt Request 1
	IRQ_INT2          = 5  // External Interrupt Request 2
	IRQ_INT3          = 6  // External Interrupt Request 3
	IRQ_PCINT0        = 7  // Pin Change Interrupt 0
	IRQ_PCINT1        = 8  // Pin Change Interrupt 1
	IRQ_WDT           = 9  // Watchdog Timeout Interrupt
	IRQ_BGSCD         = 10 // Bandgap Buffer Short Circuit Detected
	IRQ_CHDET         = 11 // Charger Detect
	IRQ_TIMER1_IC     = 12 // Timer 1 Input capture
	IRQ_TIMER1_COMPA  = 13 // Timer 1 Compare Match A
	IRQ_TIMER1_COMPB  = 14 // Timer 1 Compare Match B
	IRQ_TIMER1_OVF    = 15 // Timer 1 overflow
	IRQ_TIMER0_IC     = 16 // Timer 0 Input Capture
	IRQ_TIMER0_COMPA  = 17 // Timer 0 Comapre Match A
	IRQ_TIMER0_COMPB  = 18 // Timer 0 Compare Match B
	IRQ_TIMER0_OVF    = 19 // Timer 0 Overflow
	IRQ_TWIBUSCD      = 20 // Two-Wire Bus Connect/Disconnect
	IRQ_TWI           = 21 // Two-Wire Serial Interface
	IRQ_SPI_STC       = 22 // SPI Serial transfer complete
	IRQ_VADC          = 23 // Voltage ADC Conversion Complete
	IRQ_CCADC_CONV    = 24 // Coulomb Counter ADC Conversion Complete
	IRQ_CCADC_REG_CUR = 25 // Coloumb Counter ADC Regular Current
	IRQ_CCADC_ACC     = 26 // Coloumb Counter ADC Accumulator
	IRQ_EE_READY      = 27 // EEPROM Ready
	IRQ_SPM           = 28 // SPM Ready
	IRQ_max           = 28 // Highest interrupt number on this device.
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

//export __vector_BPINT
//go:interrupt
func interruptBPINT() {
	callHandlers(IRQ_BPINT)
}

//export __vector_VREGMON
//go:interrupt
func interruptVREGMON() {
	callHandlers(IRQ_VREGMON)
}

//export __vector_INT0
//go:interrupt
func interruptINT0() {
	callHandlers(IRQ_INT0)
}

//export __vector_INT1
//go:interrupt
func interruptINT1() {
	callHandlers(IRQ_INT1)
}

//export __vector_INT2
//go:interrupt
func interruptINT2() {
	callHandlers(IRQ_INT2)
}

//export __vector_INT3
//go:interrupt
func interruptINT3() {
	callHandlers(IRQ_INT3)
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

//export __vector_BGSCD
//go:interrupt
func interruptBGSCD() {
	callHandlers(IRQ_BGSCD)
}

//export __vector_CHDET
//go:interrupt
func interruptCHDET() {
	callHandlers(IRQ_CHDET)
}

//export __vector_TIMER1_IC
//go:interrupt
func interruptTIMER1_IC() {
	callHandlers(IRQ_TIMER1_IC)
}

//export __vector_TIMER1_COMPA
//go:interrupt
func interruptTIMER1_COMPA() {
	callHandlers(IRQ_TIMER1_COMPA)
}

//export __vector_TIMER1_COMPB
//go:interrupt
func interruptTIMER1_COMPB() {
	callHandlers(IRQ_TIMER1_COMPB)
}

//export __vector_TIMER1_OVF
//go:interrupt
func interruptTIMER1_OVF() {
	callHandlers(IRQ_TIMER1_OVF)
}

//export __vector_TIMER0_IC
//go:interrupt
func interruptTIMER0_IC() {
	callHandlers(IRQ_TIMER0_IC)
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

//export __vector_TIMER0_OVF
//go:interrupt
func interruptTIMER0_OVF() {
	callHandlers(IRQ_TIMER0_OVF)
}

//export __vector_TWIBUSCD
//go:interrupt
func interruptTWIBUSCD() {
	callHandlers(IRQ_TWIBUSCD)
}

//export __vector_TWI
//go:interrupt
func interruptTWI() {
	callHandlers(IRQ_TWI)
}

//export __vector_SPI_STC
//go:interrupt
func interruptSPI_STC() {
	callHandlers(IRQ_SPI_STC)
}

//export __vector_VADC
//go:interrupt
func interruptVADC() {
	callHandlers(IRQ_VADC)
}

//export __vector_CCADC_CONV
//go:interrupt
func interruptCCADC_CONV() {
	callHandlers(IRQ_CCADC_CONV)
}

//export __vector_CCADC_REG_CUR
//go:interrupt
func interruptCCADC_REG_CUR() {
	callHandlers(IRQ_CCADC_REG_CUR)
}

//export __vector_CCADC_ACC
//go:interrupt
func interruptCCADC_ACC() {
	callHandlers(IRQ_CCADC_ACC)
}

//export __vector_EE_READY
//go:interrupt
func interruptEE_READY() {
	callHandlers(IRQ_EE_READY)
}

//export __vector_SPM
//go:interrupt
func interruptSPM() {
	callHandlers(IRQ_SPM)
}

// Peripherals.
var (
	// Fuses
	LOW  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))
	HIGH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1)))

	// Lockbits
	LOCKBIT = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Analog-to-Digital Converter
	VADMUX = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7c)))
	VADCL  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x78)))
	VADCH  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x79)))
	VADCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7a)))

	// Watchdog Timer
	WDTCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x60)))

	// FET Control
	FCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf0)))

	// Serial Peripheral Interface
	SPCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))
	SPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	SPDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))
	EEARH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x42)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))

	// Coulomb Counter
	CADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe6)))
	CADCSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe7)))
	CADCSRC = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe8)))
	CADICL  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe4)))
	CADICH  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe5)))
	CADAC3  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe3)))
	CADAC2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe2)))
	CADAC1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe1)))
	CADAC0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe0)))
	CADRCC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe9)))
	CADRDC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xea)))

	// Two Wire Serial Interface
	TWBCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xbe)))
	TWAMR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xbd)))
	TWBR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb8)))
	TWCR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xbc)))
	TWSR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb9)))
	TWDR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xbb)))
	TWAR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xba)))

	// External Interrupts
	EICRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x69)))
	EIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))
	PCICR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x68)))
	PCIFR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3b)))
	PCMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6c)))
	PCMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6b)))

	// Timer/Counter, 16-bit
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x81)))
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x80)))
	TCNT1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x84)))
	TCNT1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x85)))
	OCR1A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x88)))
	OCR1B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x89)))
	TIMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6f)))
	TIFR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))
	TCCR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x45)))
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x44)))
	TCNT0L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	TCNT0H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x47)))
	OCR0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x48)))
	OCR0B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x49)))
	TIMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6e)))
	TIFR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))

	// Cell Balancing
	CBCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf1)))

	// Battery Protection
	BPPLR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xfe)))
	BPCR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xfd)))
	BPHCTR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xfc)))
	BPOCTR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xfb)))
	BPSCTR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xfa)))
	BPCHCD = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf9)))
	BPDHCD = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf8)))
	BPCOCD = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf7)))
	BPDOCD = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf6)))
	BPSCD  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf5)))
	BPIFR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf3)))
	BPIMSK = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf2)))

	// Charger Detect
	CHGDCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xd4)))

	// Voltage Regulator
	ROCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc8)))

	// Bandgap
	BGCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xd2)))
	BGCRR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xd1)))
	BGCCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xd0)))

	// CPU Registers
	SREG    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	SPL     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	MCUCR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x55)))
	MCUSR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x54)))
	FOSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x66)))
	OSICSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	SMCR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	GPIOR2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	GPIOR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	GPIOR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	DIDR0   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7e)))
	PRR0    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x64)))
	CLKPR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x61)))

	// I/O Port
	PORTA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x22)))
	DDRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x21)))
	PINA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x20)))
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	PORTC = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))
	PINC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))

	// Bootloader
	SPMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))
)

// Bitfields for FUSE: Fuses
const (
	// LOW
	LOW_WDTON   = 0x80 // Watch-dog Timer always on
	LOW_EESAVE  = 0x40 // Preserve EEPROM through the Chip Erase cycle
	LOW_SPIEN   = 0x20 // Serial program downloading (SPI) enabled
	LOW_SUT0    = 0x4  // Select start-up time
	LOW_SUT1    = 0x8  // Select start-up time
	LOW_SUT2    = 0x10 // Select start-up time
	LOW_OSCSEL0 = 0x1  // Oscillator select
	LOW_OSCSEL1 = 0x2  // Oscillator select

	// HIGH
	HIGH_DUVRDINIT = 0x10 // DUVR mode on
	HIGH_DWEN      = 0x8  // Debug Wire enable
	HIGH_BOOTSZ0   = 0x2  // Select Boot Size
	HIGH_BOOTSZ1   = 0x4  // Select Boot Size
	HIGH_BOOTRST   = 0x1  // Boot Reset vector Enabled
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB0   = 0x1  // Memory Lock
	LOCKBIT_LB1   = 0x2  // Memory Lock
	LOCKBIT_BLB00 = 0x4  // Boot Loader Protection Mode
	LOCKBIT_BLB01 = 0x8  // Boot Loader Protection Mode
	LOCKBIT_BLB10 = 0x10 // Boot Loader Protection Mode
	LOCKBIT_BLB11 = 0x20 // Boot Loader Protection Mode
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// VADMUX: The VADC multiplexer Selection Register
	VADMUX_VADMUX0 = 0x1 // Analog Channel and Gain Selection Bits
	VADMUX_VADMUX1 = 0x2 // Analog Channel and Gain Selection Bits
	VADMUX_VADMUX2 = 0x4 // Analog Channel and Gain Selection Bits
	VADMUX_VADMUX3 = 0x8 // Analog Channel and Gain Selection Bits

	// VADCSR: The VADC Control and Status register
	VADCSR_VADEN   = 0x8 // VADC Enable
	VADCSR_VADSC   = 0x4 // VADC Satrt Conversion
	VADCSR_VADCCIF = 0x2 // VADC Conversion Complete Interrupt Flag
	VADCSR_VADCCIE = 0x1 // VADC Conversion Complete Interrupt Enable
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

// Bitfields for FET: FET Control
const (
	// FCSR: FET Control and Status Register
	FCSR_DUVRD = 0x8 // Deep Under-Voltage Recovery Disable
	FCSR_CPS   = 0x4 // Current Protection Status
	FCSR_DFE   = 0x2 // Discharge FET Enable
	FCSR_CFE   = 0x1 // Charge FET Enable
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPCR: SPI Control Register
	SPCR_SPIE = 0x80 // SPI Interrupt Enable
	SPCR_SPE  = 0x40 // SPI Enable
	SPCR_DORD = 0x20 // Data Order
	SPCR_MSTR = 0x10 // Master/Slave Select
	SPCR_CPOL = 0x8  // Clock polarity
	SPCR_CPHA = 0x4  // Clock Phase
	SPCR_SPR0 = 0x1  // SPI Clock Rate Selects
	SPCR_SPR1 = 0x2  // SPI Clock Rate Selects

	// SPSR: SPI Status Register
	SPSR_SPIF  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL  = 0x40 // Write Collision Flag
	SPSR_SPI2X = 0x1  // Double SPI Speed Bit
)

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EEPM0 = 0x10
	EECR_EEPM1 = 0x20
	EECR_EERIE = 0x8 // EEProm Ready Interrupt Enable
	EECR_EEMPE = 0x4 // EEPROM Master Write Enable
	EECR_EEPE  = 0x2 // EEPROM Write Enable
	EECR_EERE  = 0x1 // EEPROM Read Enable
)

// Bitfields for COULOMB_COUNTER: Coulomb Counter
const (
	// CADCSRA: CC-ADC Control and Status Register A
	CADCSRA_CADEN  = 0x80 // When the CADEN bit is cleared (zero), the CC-ADC is disabled. When the CADEN bit is set (one), the CC-ADC will continuously measure the voltage drop over the external sense resistor RSENSE. In Power-down, only the Regular Current detection is active. In Power-off, the CC-ADC is always disabled.
	CADCSRA_CADPOL = 0x40
	CADCSRA_CADUB  = 0x20 // CC_ADC Update Busy
	CADCSRA_CADAS0 = 0x8  // CC_ADC Accumulate Current Select Bits
	CADCSRA_CADAS1 = 0x10 // CC_ADC Accumulate Current Select Bits
	CADCSRA_CADSI0 = 0x2  // The CADSI bits determine the current sampling interval for the Regular Current detection in Power-down mode. The actual settings remain to be determined.
	CADCSRA_CADSI1 = 0x4  // The CADSI bits determine the current sampling interval for the Regular Current detection in Power-down mode. The actual settings remain to be determined.
	CADCSRA_CADSE  = 0x1  // When the CADSE bit is written to one, the ongoing CC-ADC conversion is aborted, and the CC-ADC enters Regular Current detection mode.

	// CADCSRB: CC-ADC Control and Status Register B
	CADCSRB_CADACIE = 0x40
	CADCSRB_CADRCIE = 0x20 // Regular Current Interrupt Enable
	CADCSRB_CADICIE = 0x10 // CAD Instantenous Current Interrupt Enable
	CADCSRB_CADACIF = 0x4  // CC-ADC Accumulate Current Interrupt Flag
	CADCSRB_CADRCIF = 0x2  // CC-ADC Accumulate Current Interrupt Flag
	CADCSRB_CADICIF = 0x1  // CC-ADC Instantaneous Current Interrupt Flag

	// CADCSRC: CC-ADC Control and Status Register C
	CADCSRC_CADVSE = 0x1 // CC-ADC Voltage Scaling Enable
)

// Bitfields for TWI: Two Wire Serial Interface
const (
	// TWBCSR: TWI Bus Control and Status Register
	TWBCSR_TWBCIF = 0x80 // TWI Bus Connect/Disconnect Interrupt Flag
	TWBCSR_TWBCIE = 0x40 // TWI Bus Connect/Disconnect Interrupt Enable
	TWBCSR_TWBDT0 = 0x2  // TWI Bus Disconnect Time-out Period
	TWBCSR_TWBDT1 = 0x4  // TWI Bus Disconnect Time-out Period
	TWBCSR_TWBCIP = 0x1  // TWI Bus Connect/Disconnect Interrupt Polarity

	// TWAMR: TWI (Slave) Address Mask Register
	TWAMR_TWAM0 = 0x2
	TWAMR_TWAM1 = 0x4
	TWAMR_TWAM2 = 0x8
	TWAMR_TWAM3 = 0x10
	TWAMR_TWAM4 = 0x20
	TWAMR_TWAM5 = 0x40
	TWAMR_TWAM6 = 0x80

	// TWCR: TWI Control Register
	TWCR_TWINT = 0x80 // TWI Interrupt Flag
	TWCR_TWEA  = 0x40 // TWI Enable Acknowledge Bit
	TWCR_TWSTA = 0x20 // TWI Start Condition Bit
	TWCR_TWSTO = 0x10 // TWI Stop Condition Bit
	TWCR_TWWC  = 0x8  // TWI Write Collition Flag
	TWCR_TWEN  = 0x4  // TWI Enable Bit
	TWCR_TWIE  = 0x1  // TWI Interrupt Enable

	// TWSR: TWI Status Register
	TWSR_TWS0  = 0x8  // TWI Status
	TWSR_TWS1  = 0x10 // TWI Status
	TWSR_TWS2  = 0x20 // TWI Status
	TWSR_TWS3  = 0x40 // TWI Status
	TWSR_TWS4  = 0x80 // TWI Status
	TWSR_TWPS0 = 0x1  // TWI Prescaler
	TWSR_TWPS1 = 0x2  // TWI Prescaler

	// TWAR: TWI (Slave) Address register
	TWAR_TWA0  = 0x2  // TWI (Slave) Address register Bits
	TWAR_TWA1  = 0x4  // TWI (Slave) Address register Bits
	TWAR_TWA2  = 0x8  // TWI (Slave) Address register Bits
	TWAR_TWA3  = 0x10 // TWI (Slave) Address register Bits
	TWAR_TWA4  = 0x20 // TWI (Slave) Address register Bits
	TWAR_TWA5  = 0x40 // TWI (Slave) Address register Bits
	TWAR_TWA6  = 0x80 // TWI (Slave) Address register Bits
	TWAR_TWGCE = 0x1  // TWI General Call Recognition Enable Bit
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register
	EICRA_ISC30 = 0x40 // External Interrupt Sense Control 3 Bits
	EICRA_ISC31 = 0x80 // External Interrupt Sense Control 3 Bits
	EICRA_ISC20 = 0x10 // External Interrupt Sense Control 2 Bits
	EICRA_ISC21 = 0x20 // External Interrupt Sense Control 2 Bits
	EICRA_ISC10 = 0x4  // External Interrupt Sense Control 1 Bits
	EICRA_ISC11 = 0x8  // External Interrupt Sense Control 1 Bits
	EICRA_ISC00 = 0x1  // External Interrupt Sense Control 0 Bits
	EICRA_ISC01 = 0x2  // External Interrupt Sense Control 0 Bits

	// EIMSK: External Interrupt Mask Register
	EIMSK_INT0 = 0x1 // External Interrupt Request 3 Enable
	EIMSK_INT1 = 0x2 // External Interrupt Request 3 Enable
	EIMSK_INT2 = 0x4 // External Interrupt Request 3 Enable
	EIMSK_INT3 = 0x8 // External Interrupt Request 3 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_INTF0 = 0x1 // External Interrupt Flags
	EIFR_INTF1 = 0x2 // External Interrupt Flags
	EIFR_INTF2 = 0x4 // External Interrupt Flags
	EIFR_INTF3 = 0x8 // External Interrupt Flags

	// PCICR: Pin Change Interrupt Control Register
	PCICR_PCIE0 = 0x1 // Pin Change Interrupt Enables
	PCICR_PCIE1 = 0x2 // Pin Change Interrupt Enables

	// PCIFR: Pin Change Interrupt Flag Register
	PCIFR_PCIF0 = 0x1 // Pin Change Interrupt Flags
	PCIFR_PCIF1 = 0x2 // Pin Change Interrupt Flags
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_CS0 = 0x1 // Clock Select1 bis
	TCCR1B_CS1 = 0x2 // Clock Select1 bis
	TCCR1B_CS2 = 0x4 // Clock Select1 bis

	// TCCR1A: Timer/Counter 1 Control Register A
	TCCR1A_TCW1  = 0x80 // Timer/Counter Width
	TCCR1A_ICEN1 = 0x40 // Input Capture Mode Enable
	TCCR1A_ICNC1 = 0x20 // Input Capture Noise Canceler
	TCCR1A_ICES1 = 0x10 // Input Capture Edge Select
	TCCR1A_ICS1  = 0x8  // Input Capture Select
	TCCR1A_WGM10 = 0x1  // Waveform Generation Mode

	// TIMSK1: Timer/Counter Interrupt Mask Register
	TIMSK1_ICIE1  = 0x8 // Timer/Counter n Input Capture Interrupt Enable
	TIMSK1_OCIE1B = 0x4 // Timer/Counter1 Output Compare B Interrupt Enable
	TIMSK1_OCIE1A = 0x2 // Timer/Counter1 Output Compare A Interrupt Enable
	TIMSK1_TOIE1  = 0x1 // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter Interrupt Flag register
	TIFR1_ICF1  = 0x8 // Timer/Counter 1 Input Capture Flag
	TIFR1_OCF1B = 0x4 // Timer/Counter1 Output Compare Flag B
	TIFR1_OCF1A = 0x2 // Timer/Counter1 Output Compare Flag A
	TIFR1_TOV1  = 0x1 // Timer/Counter1 Overflow Flag

	// TCCR0B: Timer/Counter0 Control Register B
	TCCR0B_CS02 = 0x4 // Clock Select0 bit 2
	TCCR0B_CS01 = 0x2 // Clock Select0 bit 1
	TCCR0B_CS00 = 0x1 // Clock Select0 bit 0

	// TCCR0A: Timer/Counter 0 Control Register A
	TCCR0A_TCW0  = 0x80 // Timer/Counter Width
	TCCR0A_ICEN0 = 0x40 // Input Capture Mode Enable
	TCCR0A_ICNC0 = 0x20 // Input Capture Noise Canceler
	TCCR0A_ICES0 = 0x10 // Input Capture Edge Select
	TCCR0A_ICS0  = 0x8  // Input Capture Select
	TCCR0A_WGM00 = 0x1  // Waveform Generation Mode

	// TIMSK0: Timer/Counter Interrupt Mask Register
	TIMSK0_ICIE0  = 0x8 // Timer/Counter n Input Capture Interrupt Enable
	TIMSK0_OCIE0B = 0x4 // Timer/Counter0 Output Compare B Interrupt Enable
	TIMSK0_OCIE0A = 0x2 // Timer/Counter0 Output Compare A Interrupt Enable
	TIMSK0_TOIE0  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter Interrupt Flag register
	TIFR0_ICF0  = 0x8 // Timer/Counter 0 Input Capture Flag
	TIFR0_OCF0B = 0x4 // Timer/Counter0 Output Compare Flag B
	TIFR0_OCF0A = 0x2 // Timer/Counter0 Output Compare Flag A
	TIFR0_TOV0  = 0x1 // Timer/Counter0 Overflow Flag
)

// Bitfields for CELL_BALANCING: Cell Balancing
const (
	// CBCR: Cell Balancing Control Register
	CBCR_CBE0 = 0x1 // Cell Balancing Enables
	CBCR_CBE1 = 0x2 // Cell Balancing Enables
	CBCR_CBE2 = 0x4 // Cell Balancing Enables
	CBCR_CBE3 = 0x8 // Cell Balancing Enables
)

// Bitfields for BATTERY_PROTECTION: Battery Protection
const (
	// BPPLR: Battery Protection Parameter Lock Register
	BPPLR_BPPLE = 0x2 // Battery Protection Parameter Lock Enable
	BPPLR_BPPL  = 0x1 // Battery Protection Parameter Lock

	// BPCR: Battery Protection Control Register
	BPCR_EPID = 0x20 // External Protection Input Disable
	BPCR_SCD  = 0x10 // Short Circuit Protection Disabled
	BPCR_DOCD = 0x8  // Discharge Over-current Protection Disabled
	BPCR_COCD = 0x4  // Charge Over-current Protection Disabled
	BPCR_DHCD = 0x2  // Discharge High-current Protection Disable
	BPCR_CHCD = 0x1  // Charge High-current Protection Disable

	// BPIFR: Battery Protection Interrupt Flag Register
	BPIFR_SCIF  = 0x10 // Short-circuit Protection Activated Interrupt Flag
	BPIFR_DOCIF = 0x8  // Discharge Over-current Protection Activated Interrupt Flag
	BPIFR_COCIF = 0x4  // Charge Over-current Protection Activated Interrupt Flag
	BPIFR_DHCIF = 0x2  // Disharge High-current Protection Activated Interrupt
	BPIFR_CHCIF = 0x1  // Charge High-current Protection Activated Interrupt

	// BPIMSK: Battery Protection Interrupt Mask Register
	BPIMSK_SCIE  = 0x10 // Short-circuit Protection Activated Interrupt Enable
	BPIMSK_DOCIE = 0x8  // Discharge Over-current Protection Activated Interrupt Enable
	BPIMSK_COCIE = 0x4  // Charge Over-current Protection Activated Interrupt Enable
	BPIMSK_DHCIE = 0x2  // Discharger High-current Protection Activated Interrupt
	BPIMSK_CHCIE = 0x1  // Charger High-current Protection Activated Interrupt
)

// Bitfields for CHARGER_DETECT: Charger Detect
const (
	// CHGDCSR: Charger Detect Control and Status Register
	CHGDCSR_BATTPVL  = 0x10 // BATT Pin Voltage Level
	CHGDCSR_CHGDISC0 = 0x4  // Charger Detect Interrupt Sense Control
	CHGDCSR_CHGDISC1 = 0x8  // Charger Detect Interrupt Sense Control
	CHGDCSR_CHGDIF   = 0x2  // Charger Detect Interrupt Flag
	CHGDCSR_CHGDIE   = 0x1  // Charger Detect Interrupt Enable
)

// Bitfields for VOLTAGE_REGULATOR: Voltage Regulator
const (
	// ROCR: Regulator Operating Condition Register
	ROCR_ROCS   = 0x80 // ROC Status
	ROCR_ROCD   = 0x10 // ROC Disable
	ROCR_ROCWIF = 0x2  // ROC Warning Interrupt Flag
	ROCR_ROCWIE = 0x1  // ROC Warning Interrupt Enable
)

// Bitfields for BANDGAP: Bandgap
const (
	// BGCSR: Bandgap Control and Status Register
	BGCSR_BGD     = 0x20 // Bandgap Disable
	BGCSR_BGSCDE  = 0x10 // Bandgap Short Circuit Detection Enabled
	BGCSR_BGSCDIF = 0x2  // Bandgap Short Circuit Detection Interrupt Flag
	BGCSR_BGSCDIE = 0x1  // Bandgap Short Circuit Detection Interrupt Enable

	// BGCCR: Bandgap Calibration Register
	BGCCR_BGCC0 = 0x1  // BG Calibration of PTAT Current Bits
	BGCCR_BGCC1 = 0x2  // BG Calibration of PTAT Current Bits
	BGCCR_BGCC2 = 0x4  // BG Calibration of PTAT Current Bits
	BGCCR_BGCC3 = 0x8  // BG Calibration of PTAT Current Bits
	BGCCR_BGCC4 = 0x10 // BG Calibration of PTAT Current Bits
	BGCCR_BGCC5 = 0x20 // BG Calibration of PTAT Current Bits
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

	// MCUCR: MCU Control Register
	MCUCR_CKOE  = 0x20 // Clock Output Enable
	MCUCR_PUD   = 0x10 // Pull-up disable
	MCUCR_IVSEL = 0x2  // Interrupt Vector Select
	MCUCR_IVCE  = 0x1  // Interrupt Vector Change Enable

	// MCUSR: MCU Status Register
	MCUSR_OCDRF = 0x10 // OCD Reset Flag
	MCUSR_WDRF  = 0x8  // Watchdog Reset Flag
	MCUSR_BODRF = 0x4  // Brown-out Reset Flag
	MCUSR_EXTRF = 0x2  // External Reset Flag
	MCUSR_PORF  = 0x1  // Power-on reset flag

	// OSICSR: Oscillator Sampling Interface Control and Status Register
	OSICSR_OSISEL0 = 0x10 // Oscillator Sampling Interface Select 0
	OSICSR_OSIST   = 0x2  // Oscillator Sampling Interface Status
	OSICSR_OSIEN   = 0x1  // Oscillator Sampling Interface Enable

	// SMCR: Sleep Mode Control Register
	SMCR_SM0 = 0x2 // Sleep Mode Select bits
	SMCR_SM1 = 0x4 // Sleep Mode Select bits
	SMCR_SM2 = 0x8 // Sleep Mode Select bits
	SMCR_SE  = 0x1 // Sleep Enable

	// DIDR0: Digital Input Disable Register
	DIDR0_PA1DID = 0x2 // When this bit is written logic one, the digital input buffer of the corresponding V_ADC pin is disabled.
	DIDR0_PA0DID = 0x1 // When this bit is written logic one, the digital input buffer of the corresponding V_ADC pin is disabled.

	// PRR0: Power Reduction Register 0
	PRR0_PRTWI  = 0x40 // Power Reduction TWI
	PRR0_PRVRM  = 0x20 // Power Reduction Voltage Regulator Monitor
	PRR0_PRSPI  = 0x8  // Power reduction SPI
	PRR0_PRTIM1 = 0x4  // Power Reduction Timer/Counter1
	PRR0_PRTIM0 = 0x2  // Power Reduction Timer/Counter0
	PRR0_PRVADC = 0x1  // Power Reduction V-ADC

	// CLKPR: Clock Prescale Register
	CLKPR_CLKPCE = 0x80 // Clock Prescaler Change Enable
	CLKPR_CLKPS0 = 0x1  // Clock Prescaler Select Bits
	CLKPR_CLKPS1 = 0x2  // Clock Prescaler Select Bits
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control and Status Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read-While-Write Section Busy
	SPMCSR_SIGRD  = 0x20 // Signature Row Read
	SPMCSR_RWWSRE = 0x10 // Read-While-Write Section Read Enable
	SPMCSR_LBSET  = 0x8  // Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory Enable
)
