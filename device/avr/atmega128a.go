// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATmega128A.atdf, see http://packs.download.atmel.com/

//go:build avr && atmega128a
// +build avr,atmega128a

// Device information for the ATmega128A.
package avr

import (
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATmega128A"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin, Power-on Reset, Brown-out Reset, Watchdog Reset and JTAG AVR Reset
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_INT1         = 2  // External Interrupt Request 1
	IRQ_INT2         = 3  // External Interrupt Request 2
	IRQ_INT3         = 4  // External Interrupt Request 3
	IRQ_INT4         = 5  // External Interrupt Request 4
	IRQ_INT5         = 6  // External Interrupt Request 5
	IRQ_INT6         = 7  // External Interrupt Request 6
	IRQ_INT7         = 8  // External Interrupt Request 7
	IRQ_TIMER2_COMP  = 9  // Timer/Counter2 Compare Match
	IRQ_TIMER2_OVF   = 10 // Timer/Counter2 Overflow
	IRQ_TIMER1_CAPT  = 11 // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 12 // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 13 // Timer/Counter Compare Match B
	IRQ_TIMER1_OVF   = 14 // Timer/Counter1 Overflow
	IRQ_TIMER0_COMP  = 15 // Timer/Counter0 Compare Match
	IRQ_TIMER0_OVF   = 16 // Timer/Counter0 Overflow
	IRQ_SPI_STC      = 17 // SPI Serial Transfer Complete
	IRQ_USART0_RX    = 18 // USART0, Rx Complete
	IRQ_USART0_UDRE  = 19 // USART0 Data Register Empty
	IRQ_USART0_TX    = 20 // USART0, Tx Complete
	IRQ_ADC          = 21 // ADC Conversion Complete
	IRQ_EE_READY     = 22 // EEPROM Ready
	IRQ_ANALOG_COMP  = 23 // Analog Comparator
	IRQ_TIMER1_COMPC = 24 // Timer/Counter1 Compare Match C
	IRQ_TIMER3_CAPT  = 25 // Timer/Counter3 Capture Event
	IRQ_TIMER3_COMPA = 26 // Timer/Counter3 Compare Match A
	IRQ_TIMER3_COMPB = 27 // Timer/Counter3 Compare Match B
	IRQ_TIMER3_COMPC = 28 // Timer/Counter3 Compare Match C
	IRQ_TIMER3_OVF   = 29 // Timer/Counter3 Overflow
	IRQ_USART1_RX    = 30 // USART1, Rx Complete
	IRQ_USART1_UDRE  = 31 // USART1, Data Register Empty
	IRQ_USART1_TX    = 32 // USART1, Tx Complete
	IRQ_TWI          = 33 // 2-wire Serial Interface
	IRQ_SPM_READY    = 34 // Store Program Memory Read
	IRQ_max          = 34 // Highest interrupt number on this device.
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

//export __vector_TIMER2_COMP
//go:interrupt
func interruptTIMER2_COMP() {
	callHandlers(IRQ_TIMER2_COMP)
}

//export __vector_TIMER2_OVF
//go:interrupt
func interruptTIMER2_OVF() {
	callHandlers(IRQ_TIMER2_OVF)
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

//export __vector_TIMER1_OVF
//go:interrupt
func interruptTIMER1_OVF() {
	callHandlers(IRQ_TIMER1_OVF)
}

//export __vector_TIMER0_COMP
//go:interrupt
func interruptTIMER0_COMP() {
	callHandlers(IRQ_TIMER0_COMP)
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

//export __vector_USART0_RX
//go:interrupt
func interruptUSART0_RX() {
	callHandlers(IRQ_USART0_RX)
}

//export __vector_USART0_UDRE
//go:interrupt
func interruptUSART0_UDRE() {
	callHandlers(IRQ_USART0_UDRE)
}

//export __vector_USART0_TX
//go:interrupt
func interruptUSART0_TX() {
	callHandlers(IRQ_USART0_TX)
}

//export __vector_ADC
//go:interrupt
func interruptADC() {
	callHandlers(IRQ_ADC)
}

//export __vector_EE_READY
//go:interrupt
func interruptEE_READY() {
	callHandlers(IRQ_EE_READY)
}

//export __vector_ANALOG_COMP
//go:interrupt
func interruptANALOG_COMP() {
	callHandlers(IRQ_ANALOG_COMP)
}

//export __vector_TIMER1_COMPC
//go:interrupt
func interruptTIMER1_COMPC() {
	callHandlers(IRQ_TIMER1_COMPC)
}

//export __vector_TIMER3_CAPT
//go:interrupt
func interruptTIMER3_CAPT() {
	callHandlers(IRQ_TIMER3_CAPT)
}

//export __vector_TIMER3_COMPA
//go:interrupt
func interruptTIMER3_COMPA() {
	callHandlers(IRQ_TIMER3_COMPA)
}

//export __vector_TIMER3_COMPB
//go:interrupt
func interruptTIMER3_COMPB() {
	callHandlers(IRQ_TIMER3_COMPB)
}

//export __vector_TIMER3_COMPC
//go:interrupt
func interruptTIMER3_COMPC() {
	callHandlers(IRQ_TIMER3_COMPC)
}

//export __vector_TIMER3_OVF
//go:interrupt
func interruptTIMER3_OVF() {
	callHandlers(IRQ_TIMER3_OVF)
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

//export __vector_TWI
//go:interrupt
func interruptTWI() {
	callHandlers(IRQ_TWI)
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

	// Analog Comparator
	ACSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))

	// Serial Peripheral Interface
	SPDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2f)))
	SPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	SPCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))

	// Two Wire Serial Interface
	TWBR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x70)))
	TWCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x74)))
	TWSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x71)))
	TWDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x73)))
	TWAR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x72)))

	// USART
	UDR0   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))
	UCSR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	UCSR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2a)))
	UCSR0C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x95)))
	UBRR0H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x90)))
	UBRR0L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x29)))
	UDR1   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x9c)))
	UCSR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x9b)))
	UCSR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x9a)))
	UCSR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x9d)))
	UBRR1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x98)))
	UBRR1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x99)))

	// CPU Registers
	SREG   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	SPL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	MCUCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x55)))
	XMCRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6d)))
	XMCRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6c)))
	OSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6f)))
	XDIV   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5c)))
	RAMPZ  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5b)))

	// Bootloader
	SPMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x68)))

	// JTAG Interface
	OCDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x42)))

	// Other Registers

	// External Interrupts
	EICRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6a)))
	EICRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5a)))
	EIMSK = (*volatile.Register8)(unsafe.Pointer(uintptr(0x59)))
	EIFR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x58)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	EEARH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))

	// I/O Port
	PORTA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3b)))
	DDRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3a)))
	PINA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x39)))
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x38)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))
	PORTC = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))
	DDRC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x34)))
	PINC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x33)))
	PORTD = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))
	DDRD  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x31)))
	PIND  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x30)))
	PORTE = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	DDRE  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x22)))
	PINE  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x21)))
	PORTF = (*volatile.Register8)(unsafe.Pointer(uintptr(0x62)))
	DDRF  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x61)))
	PINF  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x20)))
	PORTG = (*volatile.Register8)(unsafe.Pointer(uintptr(0x65)))
	DDRG  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x64)))
	PING  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x63)))

	// Timer/Counter, 8-bit Async
	TCCR0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	TCNT0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x52)))
	OCR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x51)))
	ASSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x50)))

	// Timer/Counter, 16-bit
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4f)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))
	TCCR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7a)))
	TCNT1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))
	TCNT1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	OCR1AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	OCR1AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	OCR1BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x48)))
	OCR1BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x49)))
	OCR1CL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x78)))
	OCR1CH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x79)))
	ICR1L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	ICR1H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x47)))
	TCCR3A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8b)))
	TCCR3B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8a)))
	TCCR3C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8c)))
	TCNT3L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x88)))
	TCNT3H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x89)))
	OCR3AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x86)))
	OCR3AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x87)))
	OCR3BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x84)))
	OCR3BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x85)))
	OCR3CL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x82)))
	OCR3CH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x83)))
	ICR3L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x80)))
	ICR3H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x81)))

	// Timer/Counter, 8-bit
	TCCR2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x45)))
	TCNT2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x44)))
	OCR2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x43)))

	// Watchdog Timer
	WDTCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))

	// Analog-to-Digital Converter
	ADMUX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_M103C = 0x2 // ATmega103 Compatibility Mode
	EXTENDED_WDTON = 0x1 // Watchdog Timer always on

	// HIGH
	HIGH_OCDEN   = 0x80 // On-Chip Debug Enabled
	HIGH_JTAGEN  = 0x40 // JTAG Interface Enabled
	HIGH_SPIEN   = 0x20 // Serial program downloading (SPI) enabled
	HIGH_EESAVE  = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ0 = 0x2  // Select Boot Size
	HIGH_BOOTSZ1 = 0x4  // Select Boot Size
	HIGH_BOOTRST = 0x1  // Boot Reset vector Enabled
	HIGH_CKOPT   = 0x10 // CKOPT fuse (operation dependent of CKSEL fuses)

	// LOW
	LOW_BODLEVEL   = 0x80 // Brownout detector trigger level
	LOW_BODEN      = 0x40 // Brown-out detection enabled
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
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPSR: SPI Status Register
	SPSR_SPIF  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL  = 0x40 // Write Collision Flag
	SPSR_SPI2X = 0x1  // Double SPI Speed Bit

	// SPCR: SPI Control Register
	SPCR_SPIE = 0x80 // SPI Interrupt Enable
	SPCR_SPE  = 0x40 // SPI Enable
	SPCR_DORD = 0x20 // Data Order
	SPCR_MSTR = 0x10 // Master/Slave Select
	SPCR_CPOL = 0x8  // Clock polarity
	SPCR_CPHA = 0x4  // Clock Phase
	SPCR_SPR0 = 0x1  // SPI Clock Rate Selects
	SPCR_SPR1 = 0x2  // SPI Clock Rate Selects
)

// Bitfields for TWI: Two Wire Serial Interface
const (
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

// Bitfields for USART: USART
const (
	// UCSR0A: USART Control and Status Register A
	UCSR0A_RXC0  = 0x80 // USART Receive Complete
	UCSR0A_TXC0  = 0x40 // USART Transmitt Complete
	UCSR0A_UDRE0 = 0x20 // USART Data Register Empty
	UCSR0A_FE0   = 0x10 // Framing Error
	UCSR0A_DOR0  = 0x8  // Data overRun
	UCSR0A_UPE0  = 0x4  // Parity Error
	UCSR0A_U2X0  = 0x2  // Double the USART transmission speed
	UCSR0A_MPCM0 = 0x1  // Multi-processor Communication Mode

	// UCSR0B: USART Control and Status Register B
	UCSR0B_RXCIE0 = 0x80 // RX Complete Interrupt Enable
	UCSR0B_TXCIE0 = 0x40 // TX Complete Interrupt Enable
	UCSR0B_UDRIE0 = 0x20 // USART Data register Empty Interrupt Enable
	UCSR0B_RXEN0  = 0x10 // Receiver Enable
	UCSR0B_TXEN0  = 0x8  // Transmitter Enable
	UCSR0B_UCSZ02 = 0x4  // Character Size
	UCSR0B_RXB80  = 0x2  // Receive Data Bit 8
	UCSR0B_TXB80  = 0x1  // Transmit Data Bit 8

	// UCSR0C: USART Control and Status Register C
	UCSR0C_UMSEL0 = 0x40 // USART Mode Select
	UCSR0C_UPM00  = 0x10 // Parity Mode Bits
	UCSR0C_UPM01  = 0x20 // Parity Mode Bits
	UCSR0C_USBS0  = 0x8  // Stop Bit Select
	UCSR0C_UCSZ00 = 0x2  // Character Size
	UCSR0C_UCSZ01 = 0x4  // Character Size
	UCSR0C_UCPOL0 = 0x1  // Clock Polarity

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
	UCSR1C_UMSEL1 = 0x40 // USART Mode Select
	UCSR1C_UPM10  = 0x10 // Parity Mode Bits
	UCSR1C_UPM11  = 0x20 // Parity Mode Bits
	UCSR1C_USBS1  = 0x8  // Stop Bit Select
	UCSR1C_UCSZ10 = 0x2  // Character Size
	UCSR1C_UCSZ11 = 0x4  // Character Size
	UCSR1C_UCPOL1 = 0x1  // Clock Polarity
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
	MCUCR_SRE   = 0x80 // External SRAM Enable
	MCUCR_SRW10 = 0x40 // External SRAM Wait State Select
	MCUCR_SE    = 0x20 // Sleep Enable
	MCUCR_SM0   = 0x8  // Sleep Mode Select
	MCUCR_SM1   = 0x10 // Sleep Mode Select
	MCUCR_SM2   = 0x4  // Sleep Mode Select
	MCUCR_IVSEL = 0x2  // Interrupt Vector Select
	MCUCR_IVCE  = 0x1  // Interrupt Vector Change Enable

	// XMCRA: External Memory Control Register A
	XMCRA_SRL0  = 0x10 // Wait state page limit
	XMCRA_SRL1  = 0x20 // Wait state page limit
	XMCRA_SRL2  = 0x40 // Wait state page limit
	XMCRA_SRW00 = 0x4  // Wait state select bit lower page
	XMCRA_SRW01 = 0x8  // Wait state select bit lower page
	XMCRA_SRW11 = 0x2  // Wait state select bit upper page

	// XMCRB: External Memory Control Register B
	XMCRB_XMBK = 0x80 // External Memory Bus Keeper Enable
	XMCRB_XMM0 = 0x1  // External Memory High Mask
	XMCRB_XMM1 = 0x2  // External Memory High Mask
	XMCRB_XMM2 = 0x4  // External Memory High Mask

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL0 = 0x1  // Oscillator Calibration
	OSCCAL_OSCCAL1 = 0x2  // Oscillator Calibration
	OSCCAL_OSCCAL2 = 0x4  // Oscillator Calibration
	OSCCAL_OSCCAL3 = 0x8  // Oscillator Calibration
	OSCCAL_OSCCAL4 = 0x10 // Oscillator Calibration
	OSCCAL_OSCCAL5 = 0x20 // Oscillator Calibration
	OSCCAL_OSCCAL6 = 0x40 // Oscillator Calibration
	OSCCAL_OSCCAL7 = 0x80 // Oscillator Calibration

	// RAMPZ: RAM Page Z Select Register
	RAMPZ_RAMPZ0 = 0x1 // RAM Page Z Select Register Bit 0
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read While Write Section Busy
	SPMCSR_RWWSRE = 0x10 // Read While Write section read enable
	SPMCSR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory Enable
)

// Bitfields for JTAG: JTAG Interface
const (
	// OCDR: On-Chip Debug Related Register in I/O Memory
	OCDR_OCDR0 = 0x1  // On-Chip Debug Register Bits
	OCDR_OCDR1 = 0x2  // On-Chip Debug Register Bits
	OCDR_OCDR2 = 0x4  // On-Chip Debug Register Bits
	OCDR_OCDR3 = 0x8  // On-Chip Debug Register Bits
	OCDR_OCDR4 = 0x10 // On-Chip Debug Register Bits
	OCDR_OCDR5 = 0x20 // On-Chip Debug Register Bits
	OCDR_OCDR6 = 0x40 // On-Chip Debug Register Bits
	OCDR_OCDR7 = 0x80 // On-Chip Debug Register Bits
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
)

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EERIE = 0x8 // EEPROM Ready Interrupt Enable
	EECR_EEMWE = 0x4 // EEPROM Master Write Enable
	EECR_EEWE  = 0x2 // EEPROM Write Enable
	EECR_EERE  = 0x1 // EEPROM Read Enable
)

// Bitfields for TC8_ASYNC: Timer/Counter, 8-bit Async
const (
	// TCCR0: Timer/Counter Control Register
	TCCR0_FOC0  = 0x80 // Force Output Compare
	TCCR0_WGM00 = 0x40 // Waveform Generation Mode 0
	TCCR0_COM00 = 0x10 // Compare Match Output Modes
	TCCR0_COM01 = 0x20 // Compare Match Output Modes
	TCCR0_WGM01 = 0x8  // Waveform Generation Mode 1
	TCCR0_CS00  = 0x1  // Clock Selects
	TCCR0_CS01  = 0x2  // Clock Selects
	TCCR0_CS02  = 0x4  // Clock Selects

	// ASSR: Asynchronus Status Register
	ASSR_AS0    = 0x8 // Asynchronus Timer/Counter 0
	ASSR_TCN0UB = 0x4 // Timer/Counter0 Update Busy
	ASSR_OCR0UB = 0x2 // Output Compare register 0 Busy
	ASSR_TCR0UB = 0x1 // Timer/Counter Control Register 0 Update Busy
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
	TCCR1A_WGM10  = 0x1  // Waveform Generation Mode Bits
	TCCR1A_WGM11  = 0x2  // Waveform Generation Mode Bits

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM10 = 0x8  // Waveform Generation Mode
	TCCR1B_WGM11 = 0x10 // Waveform Generation Mode
	TCCR1B_CS10  = 0x1  // Clock Select1 bits
	TCCR1B_CS11  = 0x2  // Clock Select1 bits
	TCCR1B_CS12  = 0x4  // Clock Select1 bits

	// TCCR1C: Timer/Counter1 Control Register C
	TCCR1C_FOC1A = 0x80 // Force Output Compare for channel A
	TCCR1C_FOC1B = 0x40 // Force Output Compare for channel B
	TCCR1C_FOC1C = 0x20 // Force Output Compare for channel C

	// TCCR3A: Timer/Counter3 Control Register A
	TCCR3A_COM3A0 = 0x40 // Compare Output Mode 3A, bits
	TCCR3A_COM3A1 = 0x80 // Compare Output Mode 3A, bits
	TCCR3A_COM3B0 = 0x10 // Compare Output Mode 3B, bits
	TCCR3A_COM3B1 = 0x20 // Compare Output Mode 3B, bits
	TCCR3A_COM3C0 = 0x4  // Compare Output Mode 3C, bits
	TCCR3A_COM3C1 = 0x8  // Compare Output Mode 3C, bits
	TCCR3A_WGM30  = 0x1  // Waveform Generation Mode Bits
	TCCR3A_WGM31  = 0x2  // Waveform Generation Mode Bits

	// TCCR3B: Timer/Counter3 Control Register B
	TCCR3B_ICNC3 = 0x80 // Input Capture 3  Noise Canceler
	TCCR3B_ICES3 = 0x40 // Input Capture 3 Edge Select
	TCCR3B_WGM30 = 0x8  // Waveform Generation Mode
	TCCR3B_WGM31 = 0x10 // Waveform Generation Mode
	TCCR3B_CS30  = 0x1  // Clock Select3 bits
	TCCR3B_CS31  = 0x2  // Clock Select3 bits
	TCCR3B_CS32  = 0x4  // Clock Select3 bits

	// TCCR3C: Timer/Counter3 Control Register C
	TCCR3C_FOC3A = 0x80 // Force Output Compare for channel A
	TCCR3C_FOC3B = 0x40 // Force Output Compare for channel B
	TCCR3C_FOC3C = 0x20 // Force Output Compare for channel C
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR2: Timer/Counter Control Register
	TCCR2_FOC2  = 0x80 // Force Output Compare
	TCCR2_WGM20 = 0x40 // Wafeform Generation Mode
	TCCR2_COM20 = 0x10 // Compare Match Output Mode
	TCCR2_COM21 = 0x20 // Compare Match Output Mode
	TCCR2_WGM21 = 0x8  // Waveform Generation Mode
	TCCR2_CS20  = 0x1  // Clock Select
	TCCR2_CS21  = 0x2  // Clock Select
	TCCR2_CS22  = 0x4  // Clock Select
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCR: Watchdog Timer Control Register
	WDTCR_WDCE = 0x10 // Watchdog Change Enable
	WDTCR_WDE  = 0x8  // Watch Dog Enable
	WDTCR_WDP0 = 0x1  // Watch Dog Timer Prescaler bits
	WDTCR_WDP1 = 0x2  // Watch Dog Timer Prescaler bits
	WDTCR_WDP2 = 0x4  // Watch Dog Timer Prescaler bits
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS0 = 0x40 // Reference Selection Bits
	ADMUX_REFS1 = 0x80 // Reference Selection Bits
	ADMUX_ADLAR = 0x20 // Left Adjust Result
	ADMUX_MUX0  = 0x1  // Analog Channel and Gain Selection Bits
	ADMUX_MUX1  = 0x2  // Analog Channel and Gain Selection Bits
	ADMUX_MUX2  = 0x4  // Analog Channel and Gain Selection Bits
	ADMUX_MUX3  = 0x8  // Analog Channel and Gain Selection Bits
	ADMUX_MUX4  = 0x10 // Analog Channel and Gain Selection Bits

	// ADCSRA: The ADC Control and Status register
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADFR  = 0x20 // ADC  Free Running Select
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS0 = 0x1  // ADC  Prescaler Select Bits
	ADCSRA_ADPS1 = 0x2  // ADC  Prescaler Select Bits
	ADCSRA_ADPS2 = 0x4  // ADC  Prescaler Select Bits
)
