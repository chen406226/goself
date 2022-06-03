// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from AT90USB162.atdf, see http://packs.download.atmel.com/

//go:build avr && at90usb162
// +build avr,at90usb162

// Device information for the AT90USB162.
package avr

import (
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "AT90USB162"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin,Power-on Reset,Brown-out Reset,Watchdog Reset,and JTAG AVR Reset. See Datasheet.
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_INT1         = 2  // External Interrupt Request 1
	IRQ_INT2         = 3  // External Interrupt Request 2
	IRQ_INT3         = 4  // External Interrupt Request 3
	IRQ_INT4         = 5  // External Interrupt Request 4
	IRQ_INT5         = 6  // External Interrupt Request 5
	IRQ_INT6         = 7  // External Interrupt Request 6
	IRQ_INT7         = 8  // External Interrupt Request 7
	IRQ_PCINT0       = 9  // Pin Change Interrupt Request 0
	IRQ_PCINT1       = 10 // Pin Change Interrupt Request 1
	IRQ_USB_GEN      = 11 // USB General Interrupt Request
	IRQ_USB_COM      = 12 // USB Endpoint/Pipe Interrupt Communication Request
	IRQ_WDT          = 13 // Watchdog Time-out Interrupt
	IRQ_TIMER1_CAPT  = 14 // Timer/Counter2 Capture Event
	IRQ_TIMER1_COMPA = 15 // Timer/Counter2 Compare Match B
	IRQ_TIMER1_COMPB = 16 // Timer/Counter2 Compare Match B
	IRQ_TIMER1_COMPC = 17 // Timer/Counter2 Compare Match C
	IRQ_TIMER1_OVF   = 18 // Timer/Counter1 Overflow
	IRQ_TIMER0_COMPA = 19 // Timer/Counter0 Compare Match A
	IRQ_TIMER0_COMPB = 20 // Timer/Counter0 Compare Match B
	IRQ_TIMER0_OVF   = 21 // Timer/Counter0 Overflow
	IRQ_SPI_STC      = 22 // SPI Serial Transfer Complete
	IRQ_USART1_RX    = 23 // USART1, Rx Complete
	IRQ_USART1_UDRE  = 24 // USART1 Data register Empty
	IRQ_USART1_TX    = 25 // USART1, Tx Complete
	IRQ_ANALOG_COMP  = 26 // Analog Comparator
	IRQ_EE_READY     = 27 // EEPROM Ready
	IRQ_SPM_READY    = 28 // Store Program Memory Read
	IRQ_max          = 28 // Highest interrupt number on this device.
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

//export __vector_INT4
//go:interrupt
func interruptINT4() {
	callHandlers(IRQ_INT4)
}

//export __vector_INT5
//go:interrupt
func interruptINT5() {
	callHandlers(IRQ_INT5)
}

//export __vector_INT6
//go:interrupt
func interruptINT6() {
	callHandlers(IRQ_INT6)
}

//export __vector_INT7
//go:interrupt
func interruptINT7() {
	callHandlers(IRQ_INT7)
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

//export __vector_USB_GEN
//go:interrupt
func interruptUSB_GEN() {
	callHandlers(IRQ_USB_GEN)
}

//export __vector_USB_COM
//go:interrupt
func interruptUSB_COM() {
	callHandlers(IRQ_USB_COM)
}

//export __vector_WDT
//go:interrupt
func interruptWDT() {
	callHandlers(IRQ_WDT)
}

//export __vector_TIMER1_CAPT
//go:interrupt
func interruptTIMER1_CAPT() {
	callHandlers(IRQ_TIMER1_CAPT)
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

//export __vector_TIMER1_COMPC
//go:interrupt
func interruptTIMER1_COMPC() {
	callHandlers(IRQ_TIMER1_COMPC)
}

//export __vector_TIMER1_OVF
//go:interrupt
func interruptTIMER1_OVF() {
	callHandlers(IRQ_TIMER1_OVF)
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

//export __vector_SPI_STC
//go:interrupt
func interruptSPI_STC() {
	callHandlers(IRQ_SPI_STC)
}

//export __vector_USART1_RX
//go:interrupt
func interruptUSART1_RX() {
	callHandlers(IRQ_USART1_RX)
}

//export __vector_USART1_UDRE
//go:interrupt
func interruptUSART1_UDRE() {
	callHandlers(IRQ_USART1_UDRE)
}

//export __vector_USART1_TX
//go:interrupt
func interruptUSART1_TX() {
	callHandlers(IRQ_USART1_TX)
}

//export __vector_ANALOG_COMP
//go:interrupt
func interruptANALOG_COMP() {
	callHandlers(IRQ_ANALOG_COMP)
}

//export __vector_EE_READY
//go:interrupt
func interruptEE_READY() {
	callHandlers(IRQ_EE_READY)
}

//export __vector_SPM_READY
//go:interrupt
func interruptSPM_READY() {
	callHandlers(IRQ_SPM_READY)
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
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	PORTD = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	DDRD  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2a)))
	PIND  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x29)))
	PORTC = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))
	DDRC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	PINC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))

	// Serial Peripheral Interface
	SPCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))
	SPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	SPDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))

	// Bootloader
	SPMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))
	EEARH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x42)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))

	// Timer/Counter, 8-bit
	OCR0B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x48)))
	OCR0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x47)))
	TCNT0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	TCCR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x45)))
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x44)))
	TIMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6e)))
	TIFR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))
	GTCCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x43)))

	// Timer/Counter, 16-bit
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x80)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x81)))
	TCCR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x82)))
	TCNT1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x84)))
	TCNT1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x85)))
	OCR1AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x88)))
	OCR1AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x89)))
	OCR1BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8a)))
	OCR1BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8b)))
	OCR1CL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8c)))
	OCR1CH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8d)))
	ICR1L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x86)))
	ICR1H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x87)))
	TIMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6f)))
	TIFR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))

	// Phase Locked Loop
	PLLCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x49)))

	// USB Device Registers
	UEINT   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf4)))
	UEBCLX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf2)))
	UEDATX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf1)))
	UEIENX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf0)))
	UESTA1X = (*volatile.Register8)(unsafe.Pointer(uintptr(0xef)))
	UESTA0X = (*volatile.Register8)(unsafe.Pointer(uintptr(0xee)))
	UECFG1X = (*volatile.Register8)(unsafe.Pointer(uintptr(0xed)))
	UECFG0X = (*volatile.Register8)(unsafe.Pointer(uintptr(0xec)))
	UECONX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xeb)))
	UERST   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xea)))
	UENUM   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe9)))
	UEINTX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe8)))
	UDMFN   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe6)))
	UDFNUML = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe4)))
	UDFNUMH = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe5)))
	UDADDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe3)))
	UDIEN   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe2)))
	UDINT   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe1)))
	UDCON   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xe0)))
	USBCON  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xd8)))
	REGCR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x63)))

	// PS/2 Controller
	UPOE   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xfb)))
	PS2CON = (*volatile.Register8)(unsafe.Pointer(uintptr(0xfa)))

	// CPU Registers
	SREG    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	SPL     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	MCUCR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x55)))
	MCUSR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x54)))
	OSCCAL  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x66)))
	CLKPR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x61)))
	SMCR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	EIND    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5c)))
	GPIOR2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	GPIOR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	GPIOR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	PRR1    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x65)))
	PRR0    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x64)))
	CLKSTA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xd2)))
	CLKSEL1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xd1)))
	CLKSEL0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xd0)))
	DWDR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x51)))

	// External Interrupts
	EICRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x69)))
	EICRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6a)))
	EIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))
	PCMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6b)))
	PCMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6c)))
	PCIFR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3b)))
	PCICR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x68)))

	// USART
	UDR1   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xce)))
	UCSR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc8)))
	UCSR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc9)))
	UCSR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0xca)))
	UCSR1D = (*volatile.Register8)(unsafe.Pointer(uintptr(0xcb)))
	UBRR1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0xcc)))
	UBRR1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0xcd)))

	// Watchdog Timer
	WDTCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x60)))
	WDTCKD = (*volatile.Register8)(unsafe.Pointer(uintptr(0x62)))

	// Analog Comparator
	ACSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x50)))
	DIDR1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7f)))
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_BODLEVEL0 = 0x1 // Brown-out Detector trigger level
	EXTENDED_BODLEVEL1 = 0x2 // Brown-out Detector trigger level
	EXTENDED_BODLEVEL2 = 0x4 // Brown-out Detector trigger level
	EXTENDED_HWBE      = 0x8 // Hardware Boot Enable

	// HIGH
	HIGH_DWEN     = 0x80 // Debug Wire enable
	HIGH_RSTDISBL = 0x40 // Reset Disabled (Enable PC6 as i/o pin)
	HIGH_SPIEN    = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON    = 0x10 // Watchdog timer always on
	HIGH_EESAVE   = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ0  = 0x2  // Select Boot Size
	HIGH_BOOTSZ1  = 0x4  // Select Boot Size
	HIGH_BOOTRST  = 0x1  // Boot Reset vector Enabled

	// LOW
	LOW_CKDIV8     = 0x80 // Divide clock by 8 internally
	LOW_CKOUT      = 0x40 // Clock output on PORTC7
	LOW_SUT_CKSEL0 = 0x1  // Select Clock Source
	LOW_SUT_CKSEL1 = 0x2  // Select Clock Source
	LOW_SUT_CKSEL2 = 0x4  // Select Clock Source
	LOW_SUT_CKSEL3 = 0x8  // Select Clock Source
	LOW_SUT_CKSEL4 = 0x10 // Select Clock Source
	LOW_SUT_CKSEL5 = 0x20 // Select Clock Source
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

// Bitfields for PORT: I/O Port
const (
	// PORTC: Port C Data Register
	PORTC_PORTC0 = 0x10 // Port C Data Register bits
	PORTC_PORTC1 = 0x20 // Port C Data Register bits
	PORTC_PORTC2 = 0x40 // Port C Data Register bits
	PORTC_PORTC3 = 0x80 // Port C Data Register bits
	PORTC_PORTC0 = 0x1  // Port C Data Register bits
	PORTC_PORTC1 = 0x2  // Port C Data Register bits
	PORTC_PORTC2 = 0x4  // Port C Data Register bits

	// DDRC: Port C Data Direction Register
	DDRC_DDC0 = 0x10 // Port C Data Direction Register bits
	DDRC_DDC1 = 0x20 // Port C Data Direction Register bits
	DDRC_DDC2 = 0x40 // Port C Data Direction Register bits
	DDRC_DDC3 = 0x80 // Port C Data Direction Register bits
	DDRC_DDC0 = 0x1  // Port C Data Direction Register bits
	DDRC_DDC1 = 0x2  // Port C Data Direction Register bits
	DDRC_DDC2 = 0x4  // Port C Data Direction Register bits

	// PINC: Port C Input Pins
	PINC_PINC0 = 0x10 // Port C Input Pins bits
	PINC_PINC1 = 0x20 // Port C Input Pins bits
	PINC_PINC2 = 0x40 // Port C Input Pins bits
	PINC_PINC3 = 0x80 // Port C Input Pins bits
	PINC_PINC0 = 0x1  // Port C Input Pins bits
	PINC_PINC1 = 0x2  // Port C Input Pins bits
	PINC_PINC2 = 0x4  // Port C Input Pins bits
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

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read While Write Section Busy
	SPMCSR_SIGRD  = 0x20 // Signature Row Read
	SPMCSR_RWWSRE = 0x10 // Read While Write section read enable
	SPMCSR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory Enable
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

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR0B: Timer/Counter Control Register B
	TCCR0B_FOC0A = 0x80 // Force Output Compare A
	TCCR0B_FOC0B = 0x40 // Force Output Compare B
	TCCR0B_WGM02 = 0x8
	TCCR0B_CS00  = 0x1 // Clock Select
	TCCR0B_CS01  = 0x2 // Clock Select
	TCCR0B_CS02  = 0x4 // Clock Select

	// TCCR0A: Timer/Counter  Control Register A
	TCCR0A_COM0A0 = 0x40 // Compare Output Mode, Phase Correct PWM Mode
	TCCR0A_COM0A1 = 0x80 // Compare Output Mode, Phase Correct PWM Mode
	TCCR0A_COM0B0 = 0x10 // Compare Output Mode, Fast PWm
	TCCR0A_COM0B1 = 0x20 // Compare Output Mode, Fast PWm
	TCCR0A_WGM00  = 0x1  // Waveform Generation Mode
	TCCR0A_WGM01  = 0x2  // Waveform Generation Mode

	// TIMSK0: Timer/Counter0 Interrupt Mask Register
	TIMSK0_OCIE0B = 0x4 // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK0_OCIE0A = 0x2 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK0_TOIE0  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter0 Interrupt Flag register
	TIFR0_OCF0B = 0x4 // Timer/Counter0 Output Compare Flag 0B
	TIFR0_OCF0A = 0x2 // Timer/Counter0 Output Compare Flag 0A
	TIFR0_TOV0  = 0x1 // Timer/Counter0 Overflow Flag

	// GTCCR: General Timer/Counter Control Register
	GTCCR_TSM     = 0x80 // Timer/Counter Synchronization Mode
	GTCCR_PSRSYNC = 0x1  // Prescaler Reset Timer/Counter1 and Timer/Counter0
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A0 = 0x40 // Compare Output Mode 1A, bits
	TCCR1A_COM1A1 = 0x80 // Compare Output Mode 1A, bits
	TCCR1A_COM1B0 = 0x10 // Compare Output Mode 1B, bits
	TCCR1A_COM1B1 = 0x20 // Compare Output Mode 1B, bits
	TCCR1A_COM1C0 = 0x4  // Compare Output Mode 1C, bits
	TCCR1A_COM1C1 = 0x8  // Compare Output Mode 1C, bits
	TCCR1A_WGM10  = 0x1  // Waveform Generation Mode
	TCCR1A_WGM11  = 0x2  // Waveform Generation Mode

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM10 = 0x8  // Waveform Generation Mode
	TCCR1B_WGM11 = 0x10 // Waveform Generation Mode
	TCCR1B_CS10  = 0x1  // Prescaler source of Timer/Counter 1
	TCCR1B_CS11  = 0x2  // Prescaler source of Timer/Counter 1
	TCCR1B_CS12  = 0x4  // Prescaler source of Timer/Counter 1

	// TCCR1C: Timer/Counter 1 Control Register C
	TCCR1C_FOC1A = 0x80 // Force Output Compare 1A
	TCCR1C_FOC1B = 0x40 // Force Output Compare 1B
	TCCR1C_FOC1C = 0x20 // Force Output Compare 1C

	// TIMSK1: Timer/Counter1 Interrupt Mask Register
	TIMSK1_ICIE1  = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_OCIE1C = 0x8  // Timer/Counter1 Output Compare C Match Interrupt Enable
	TIMSK1_OCIE1B = 0x4  // Timer/Counter1 Output Compare B Match Interrupt Enable
	TIMSK1_OCIE1A = 0x2  // Timer/Counter1 Output Compare A Match Interrupt Enable
	TIMSK1_TOIE1  = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter1 Interrupt Flag register
	TIFR1_ICF1  = 0x20 // Input Capture Flag 1
	TIFR1_OCF1C = 0x8  // Output Compare Flag 1C
	TIFR1_OCF1B = 0x4  // Output Compare Flag 1B
	TIFR1_OCF1A = 0x2  // Output Compare Flag 1A
	TIFR1_TOV1  = 0x1  // Timer/Counter1 Overflow Flag
)

// Bitfields for PLL: Phase Locked Loop
const (
	// PLLCSR: PLL Status and Control register
	PLLCSR_PLLP0 = 0x4  // PLL prescaler Bits
	PLLCSR_PLLP1 = 0x8  // PLL prescaler Bits
	PLLCSR_PLLP2 = 0x10 // PLL prescaler Bits
	PLLCSR_PLLE  = 0x2  // PLL Enable Bit
	PLLCSR_PLOCK = 0x1  // PLL Lock Status Bit
)

// Bitfields for USB_DEVICE: USB Device Registers
const (
	// UEIENX
	UEIENX_FLERRE   = 0x80
	UEIENX_NAKINE   = 0x40
	UEIENX_NAKOUTE  = 0x10
	UEIENX_RXSTPE   = 0x8
	UEIENX_RXOUTE   = 0x4
	UEIENX_STALLEDE = 0x2
	UEIENX_TXINE    = 0x1

	// UESTA1X
	UESTA1X_CTRLDIR = 0x4
	UESTA1X_CURRBK0 = 0x1
	UESTA1X_CURRBK1 = 0x2

	// UESTA0X
	UESTA0X_CFGOK    = 0x80
	UESTA0X_OVERFI   = 0x40
	UESTA0X_UNDERFI  = 0x20
	UESTA0X_DTSEQ0   = 0x4
	UESTA0X_DTSEQ1   = 0x8
	UESTA0X_NBUSYBK0 = 0x1
	UESTA0X_NBUSYBK1 = 0x2

	// UECFG1X
	UECFG1X_EPSIZE0 = 0x10
	UECFG1X_EPSIZE1 = 0x20
	UECFG1X_EPSIZE2 = 0x40
	UECFG1X_EPBK0   = 0x4
	UECFG1X_EPBK1   = 0x8
	UECFG1X_ALLOC   = 0x2

	// UECFG0X
	UECFG0X_EPTYPE0 = 0x40
	UECFG0X_EPTYPE1 = 0x80
	UECFG0X_EPDIR   = 0x1

	// UECONX
	UECONX_STALLRQ  = 0x20
	UECONX_STALLRQC = 0x10
	UECONX_RSTDT    = 0x8
	UECONX_EPEN     = 0x1

	// UERST
	UERST_EPRST0 = 0x1
	UERST_EPRST1 = 0x2
	UERST_EPRST2 = 0x4
	UERST_EPRST3 = 0x8
	UERST_EPRST4 = 0x10

	// UEINTX
	UEINTX_FIFOCON  = 0x80
	UEINTX_NAKINI   = 0x40
	UEINTX_RWAL     = 0x20
	UEINTX_NAKOUTI  = 0x10
	UEINTX_RXSTPI   = 0x8
	UEINTX_RXOUTI   = 0x4
	UEINTX_STALLEDI = 0x2
	UEINTX_TXINI    = 0x1

	// UDMFN
	UDMFN_FNCERR = 0x10

	// UDADDR
	UDADDR_ADDEN = 0x80
	UDADDR_UADD0 = 0x1
	UDADDR_UADD1 = 0x2
	UDADDR_UADD2 = 0x4
	UDADDR_UADD3 = 0x8
	UDADDR_UADD4 = 0x10
	UDADDR_UADD5 = 0x20
	UDADDR_UADD6 = 0x40

	// UDIEN
	UDIEN_UPRSME  = 0x40
	UDIEN_EORSME  = 0x20
	UDIEN_WAKEUPE = 0x10
	UDIEN_EORSTE  = 0x8
	UDIEN_SOFE    = 0x4
	UDIEN_SUSPE   = 0x1

	// UDINT
	UDINT_UPRSMI  = 0x40
	UDINT_EORSMI  = 0x20
	UDINT_WAKEUPI = 0x10
	UDINT_EORSTI  = 0x8
	UDINT_SOFI    = 0x4
	UDINT_SUSPI   = 0x1

	// UDCON
	UDCON_RSTCPU = 0x4
	UDCON_RMWKUP = 0x2
	UDCON_DETACH = 0x1

	// USBCON: USB General Control Register
	USBCON_USBE   = 0x80
	USBCON_FRZCLK = 0x20

	// REGCR: Regulator Control Register
	REGCR_REGDIS = 0x1
)

// Bitfields for PS2: PS/2 Controller
const (
	// UPOE
	UPOE_UPWE0  = 0x40
	UPOE_UPWE1  = 0x80
	UPOE_UPDRV0 = 0x10
	UPOE_UPDRV1 = 0x20
	UPOE_SCKI   = 0x8
	UPOE_DATAI  = 0x4
	UPOE_DPI    = 0x2
	UPOE_DMI    = 0x1

	// PS2CON: PS2 Pad Enable register
	PS2CON_PS2EN = 0x1 // Enable
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
	MCUCR_PUD   = 0x10 // Pull-up disable
	MCUCR_IVSEL = 0x2  // Interrupt Vector Select
	MCUCR_IVCE  = 0x1  // Interrupt Vector Change Enable

	// MCUSR: MCU Status Register
	MCUSR_USBRF = 0x20 // USB reset flag
	MCUSR_WDRF  = 0x8  // Watchdog Reset Flag
	MCUSR_BORF  = 0x4  // Brown-out Reset Flag
	MCUSR_EXTRF = 0x2  // External Reset Flag
	MCUSR_PORF  = 0x1  // Power-on reset flag

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL0 = 0x1  // Oscillator Calibration
	OSCCAL_OSCCAL1 = 0x2  // Oscillator Calibration
	OSCCAL_OSCCAL2 = 0x4  // Oscillator Calibration
	OSCCAL_OSCCAL3 = 0x8  // Oscillator Calibration
	OSCCAL_OSCCAL4 = 0x10 // Oscillator Calibration
	OSCCAL_OSCCAL5 = 0x20 // Oscillator Calibration
	OSCCAL_OSCCAL6 = 0x40 // Oscillator Calibration
	OSCCAL_OSCCAL7 = 0x80 // Oscillator Calibration

	// CLKPR
	CLKPR_CLKPCE = 0x80
	CLKPR_CLKPS0 = 0x1
	CLKPR_CLKPS1 = 0x2
	CLKPR_CLKPS2 = 0x4
	CLKPR_CLKPS3 = 0x8

	// SMCR: Sleep Mode Control Register
	SMCR_SM0 = 0x2 // Sleep Mode Select bits
	SMCR_SM1 = 0x4 // Sleep Mode Select bits
	SMCR_SM2 = 0x8 // Sleep Mode Select bits
	SMCR_SE  = 0x1 // Sleep Enable

	// GPIOR2: General Purpose IO Register 2
	GPIOR2_GPIOR0 = 0x1  // General Purpose IO Register 2 bis
	GPIOR2_GPIOR1 = 0x2  // General Purpose IO Register 2 bis
	GPIOR2_GPIOR2 = 0x4  // General Purpose IO Register 2 bis
	GPIOR2_GPIOR3 = 0x8  // General Purpose IO Register 2 bis
	GPIOR2_GPIOR4 = 0x10 // General Purpose IO Register 2 bis
	GPIOR2_GPIOR5 = 0x20 // General Purpose IO Register 2 bis
	GPIOR2_GPIOR6 = 0x40 // General Purpose IO Register 2 bis
	GPIOR2_GPIOR7 = 0x80 // General Purpose IO Register 2 bis

	// GPIOR1: General Purpose IO Register 1
	GPIOR1_GPIOR0 = 0x1  // General Purpose IO Register 1 bis
	GPIOR1_GPIOR1 = 0x2  // General Purpose IO Register 1 bis
	GPIOR1_GPIOR2 = 0x4  // General Purpose IO Register 1 bis
	GPIOR1_GPIOR3 = 0x8  // General Purpose IO Register 1 bis
	GPIOR1_GPIOR4 = 0x10 // General Purpose IO Register 1 bis
	GPIOR1_GPIOR5 = 0x20 // General Purpose IO Register 1 bis
	GPIOR1_GPIOR6 = 0x40 // General Purpose IO Register 1 bis
	GPIOR1_GPIOR7 = 0x80 // General Purpose IO Register 1 bis

	// GPIOR0: General Purpose IO Register 0
	GPIOR0_GPIOR07 = 0x80 // General Purpose IO Register 0 bit 7
	GPIOR0_GPIOR06 = 0x40 // General Purpose IO Register 0 bit 6
	GPIOR0_GPIOR05 = 0x20 // General Purpose IO Register 0 bit 5
	GPIOR0_GPIOR04 = 0x10 // General Purpose IO Register 0 bit 4
	GPIOR0_GPIOR03 = 0x8  // General Purpose IO Register 0 bit 3
	GPIOR0_GPIOR02 = 0x4  // General Purpose IO Register 0 bit 2
	GPIOR0_GPIOR01 = 0x2  // General Purpose IO Register 0 bit 1
	GPIOR0_GPIOR00 = 0x1  // General Purpose IO Register 0 bit 0

	// PRR1: Power Reduction Register1
	PRR1_PRUSB    = 0x80 // Power Reduction USB
	PRR1_PRUSART1 = 0x1  // Power Reduction USART1

	// PRR0: Power Reduction Register0
	PRR0_PRTIM0 = 0x20 // Power Reduction Timer/Counter0
	PRR0_PRTIM1 = 0x8  // Power Reduction Timer/Counter1
	PRR0_PRSPI  = 0x4  // Power Reduction Serial Peripheral Interface

	// CLKSTA
	CLKSTA_RCON  = 0x2
	CLKSTA_EXTON = 0x1

	// CLKSEL1
	CLKSEL1_RCCKSEL0 = 0x10
	CLKSEL1_RCCKSEL1 = 0x20
	CLKSEL1_RCCKSEL2 = 0x40
	CLKSEL1_RCCKSEL3 = 0x80
	CLKSEL1_EXCKSEL0 = 0x1
	CLKSEL1_EXCKSEL1 = 0x2
	CLKSEL1_EXCKSEL2 = 0x4
	CLKSEL1_EXCKSEL3 = 0x8

	// CLKSEL0
	CLKSEL0_RCSUT0 = 0x40
	CLKSEL0_RCSUT1 = 0x80
	CLKSEL0_EXSUT0 = 0x10
	CLKSEL0_EXSUT1 = 0x20
	CLKSEL0_RCE    = 0x8
	CLKSEL0_EXTE   = 0x4
	CLKSEL0_CLKS   = 0x1
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register A
	EICRA_ISC30 = 0x40 // External Interrupt Sense Control Bit
	EICRA_ISC31 = 0x80 // External Interrupt Sense Control Bit
	EICRA_ISC20 = 0x10 // External Interrupt Sense Control Bit
	EICRA_ISC21 = 0x20 // External Interrupt Sense Control Bit
	EICRA_ISC10 = 0x4  // External Interrupt Sense Control Bit
	EICRA_ISC11 = 0x8  // External Interrupt Sense Control Bit
	EICRA_ISC00 = 0x1  // External Interrupt Sense Control Bit
	EICRA_ISC01 = 0x2  // External Interrupt Sense Control Bit

	// EICRB: External Interrupt Control Register B
	EICRB_ISC70 = 0x40 // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC71 = 0x80 // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC60 = 0x10 // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC61 = 0x20 // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC50 = 0x4  // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC51 = 0x8  // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC40 = 0x1  // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC41 = 0x2  // External Interrupt 7-4 Sense Control Bit

	// EIMSK: External Interrupt Mask Register
	EIMSK_INT0 = 0x1  // External Interrupt Request 7 Enable
	EIMSK_INT1 = 0x2  // External Interrupt Request 7 Enable
	EIMSK_INT2 = 0x4  // External Interrupt Request 7 Enable
	EIMSK_INT3 = 0x8  // External Interrupt Request 7 Enable
	EIMSK_INT4 = 0x10 // External Interrupt Request 7 Enable
	EIMSK_INT5 = 0x20 // External Interrupt Request 7 Enable
	EIMSK_INT6 = 0x40 // External Interrupt Request 7 Enable
	EIMSK_INT7 = 0x80 // External Interrupt Request 7 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_INTF0 = 0x1  // External Interrupt Flags
	EIFR_INTF1 = 0x2  // External Interrupt Flags
	EIFR_INTF2 = 0x4  // External Interrupt Flags
	EIFR_INTF3 = 0x8  // External Interrupt Flags
	EIFR_INTF4 = 0x10 // External Interrupt Flags
	EIFR_INTF5 = 0x20 // External Interrupt Flags
	EIFR_INTF6 = 0x40 // External Interrupt Flags
	EIFR_INTF7 = 0x80 // External Interrupt Flags

	// PCMSK0: Pin Change Mask Register 0
	PCMSK0_PCINT0 = 0x1  // Pin Change Enable Masks
	PCMSK0_PCINT1 = 0x2  // Pin Change Enable Masks
	PCMSK0_PCINT2 = 0x4  // Pin Change Enable Masks
	PCMSK0_PCINT3 = 0x8  // Pin Change Enable Masks
	PCMSK0_PCINT4 = 0x10 // Pin Change Enable Masks
	PCMSK0_PCINT5 = 0x20 // Pin Change Enable Masks
	PCMSK0_PCINT6 = 0x40 // Pin Change Enable Masks
	PCMSK0_PCINT7 = 0x80 // Pin Change Enable Masks

	// PCMSK1: Pin Change Mask Register 1
	PCMSK1_PCINT0 = 0x1
	PCMSK1_PCINT1 = 0x2
	PCMSK1_PCINT2 = 0x4
	PCMSK1_PCINT3 = 0x8
	PCMSK1_PCINT4 = 0x10

	// PCIFR: Pin Change Interrupt Flag Register
	PCIFR_PCIF0 = 0x1 // Pin Change Interrupt Flags
	PCIFR_PCIF1 = 0x2 // Pin Change Interrupt Flags

	// PCICR: Pin Change Interrupt Control Register
	PCICR_PCIE0 = 0x1 // Pin Change Interrupt Enables
	PCICR_PCIE1 = 0x2 // Pin Change Interrupt Enables
)

// Bitfields for USART: USART
const (
	// UCSR1A: USART Control and Status Register A
	UCSR1A_RXC1  = 0x80 // USART Receive Complete
	UCSR1A_TXC1  = 0x40 // USART Transmitt Complete
	UCSR1A_UDRE1 = 0x20 // USART Data Register Empty
	UCSR1A_FE1   = 0x10 // Framing Error
	UCSR1A_DOR1  = 0x8  // Data overRun
	UCSR1A_UPE1  = 0x4  // Parity Error
	UCSR1A_U2X1  = 0x2  // Double the USART transmission speed
	UCSR1A_MPCM1 = 0x1  // Multi-processor Communication Mode

	// UCSR1B: USART Control and Status Register B
	UCSR1B_RXCIE1 = 0x80 // RX Complete Interrupt Enable
	UCSR1B_TXCIE1 = 0x40 // TX Complete Interrupt Enable
	UCSR1B_UDRIE1 = 0x20 // USART Data register Empty Interrupt Enable
	UCSR1B_RXEN1  = 0x10 // Receiver Enable
	UCSR1B_TXEN1  = 0x8  // Transmitter Enable
	UCSR1B_UCSZ12 = 0x4  // Character Size
	UCSR1B_RXB81  = 0x2  // Receive Data Bit 8
	UCSR1B_TXB81  = 0x1  // Transmit Data Bit 8

	// UCSR1C: USART Control and Status Register C
	UCSR1C_UMSEL10 = 0x40 // USART Mode Select
	UCSR1C_UMSEL11 = 0x80 // USART Mode Select
	UCSR1C_UPM10   = 0x10 // Parity Mode Bits
	UCSR1C_UPM11   = 0x20 // Parity Mode Bits
	UCSR1C_USBS1   = 0x8  // Stop Bit Select
	UCSR1C_UCSZ10  = 0x2  // Character Size
	UCSR1C_UCSZ11  = 0x4  // Character Size
	UCSR1C_UCPOL1  = 0x1  // Clock Polarity

	// UCSR1D: USART Control and Status Register D
	UCSR1D_CTSEN = 0x2 // CTS Enable
	UCSR1D_RTSEN = 0x1 // RTS Enable
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

	// WDTCKD: Watchdog Timer Clock Divider
	WDTCKD_WDEWIF = 0x8 // Watchdog Early Warning Interrupt Flag
	WDTCKD_WDEWIE = 0x4 // Watchdog Early Warning Interrupt Enable
	WDTCKD_WCLKD0 = 0x1 // Watchdog Timer Clock Dividers
	WDTCKD_WCLKD1 = 0x2 // Watchdog Timer Clock Dividers
)

// Bitfields for AC: Analog Comparator
const (
	// ACSR: Analog Comparator Control And Status Register
	ACSR_ACD   = 0x80 // Analog Comparator Disable
	ACSR_ACBG  = 0x40 // Analog Comparator Bandgap Select
	ACSR_ACO   = 0x20 // Analog Compare Output
	ACSR_ACI   = 0x10 // Analog Comparator Interrupt Flag
	ACSR_ACIE  = 0x8  // Analog Comparator Interrupt Enable
	ACSR_ACIC  = 0x4  // Analog Comparator Input Capture Enable
	ACSR_ACIS0 = 0x1  // Analog Comparator Interrupt Mode Select bits
	ACSR_ACIS1 = 0x2  // Analog Comparator Interrupt Mode Select bits

	// DIDR1
	DIDR1_AIN1D = 0x2 // AIN1 Digital Input Disable
	DIDR1_AIN0D = 0x1 // AIN0 Digital Input Disable
)
