// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATmega8A.atdf, see http://packs.download.atmel.com/

//go:build avr && atmega8a
// +build avr,atmega8a

// Device information for the ATmega8A.
package avr

import (
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATmega8A"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin, Power-on Reset, Brown-out Reset and Watchdog Reset
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_INT1         = 2  // External Interrupt Request 1
	IRQ_TIMER2_COMP  = 3  // Timer/Counter2 Compare Match
	IRQ_TIMER2_OVF   = 4  // Timer/Counter2 Overflow
	IRQ_TIMER1_CAPT  = 5  // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 6  // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 7  // Timer/Counter1 Compare Match B
	IRQ_TIMER1_OVF   = 8  // Timer/Counter1 Overflow
	IRQ_TIMER0_OVF   = 9  // Timer/Counter0 Overflow
	IRQ_SPI_STC      = 10 // Serial Transfer Complete
	IRQ_USART_RXC    = 11 // USART, Rx Complete
	IRQ_USART_UDRE   = 12 // USART Data Register Empty
	IRQ_USART_TXC    = 13 // USART, Tx Complete
	IRQ_ADC          = 14 // ADC Conversion Complete
	IRQ_EE_RDY       = 15 // EEPROM Ready
	IRQ_ANA_COMP     = 16 // Analog Comparator
	IRQ_TWI          = 17 // 2-wire Serial Interface
	IRQ_SPM_RDY      = 18 // Store Program Memory Ready
	IRQ_max          = 18 // Highest interrupt number on this device.
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

//export __vector_USART_RXC
//go:interrupt
func interruptUSART_RXC() {
	callHandlers(IRQ_USART_RXC)
}

//export __vector_USART_UDRE
//go:interrupt
func interruptUSART_UDRE() {
	callHandlers(IRQ_USART_UDRE)
}

//export __vector_USART_TXC
//go:interrupt
func interruptUSART_TXC() {
	callHandlers(IRQ_USART_TXC)
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

//export __vector_ANA_COMP
//go:interrupt
func interruptANA_COMP() {
	callHandlers(IRQ_ANA_COMP)
}

//export __vector_TWI
//go:interrupt
func interruptTWI() {
	callHandlers(IRQ_TWI)
}

//export __vector_SPM_RDY
//go:interrupt
func interruptSPM_RDY() {
	callHandlers(IRQ_SPM_RDY)
}

// Peripherals.
var (
	// Fuses
	HIGH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1)))
	LOW  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Lockbits
	LOCKBIT = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Analog Comparator
	ACSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))

	// Serial Peripheral Interface
	SPDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2f)))
	SPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	SPCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))

	// External Interrupts
	GICR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5b)))
	GIFR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5a)))

	// Timer/Counter, 8-bit
	TCCR0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	TCNT0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x52)))

	// Timer/Counter, 16-bit
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4f)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))
	TCNT1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))
	TCNT1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	OCR1AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	OCR1AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	OCR1BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x48)))
	OCR1BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x49)))
	ICR1L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	ICR1H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x47)))

	// Timer/Counter, 8-bit Async
	TCCR2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x45)))
	TCNT2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x44)))
	OCR2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x43)))
	ASSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x42)))

	// USART
	UDR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))
	UCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	UCSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2a)))
	UCSRC = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	UBRRH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	UBRRL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x29)))

	// Two Wire Serial Interface
	TWBR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x20)))
	TWCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x56)))
	TWSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x21)))
	TWDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	TWAR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x22)))

	// Watchdog Timer
	WDTCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))

	// I/O Port
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x38)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))
	PORTC = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))
	DDRC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x34)))
	PINC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x33)))
	PORTD = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))
	DDRD  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x31)))
	PIND  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x30)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	EEARH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))

	// CPU Registers
	SREG   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	SPL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	MCUCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x54)))
	OSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x51)))
	SPMCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))

	// Analog-to-Digital Converter
	ADMUX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
)

// Bitfields for FUSE: Fuses
const (
	// HIGH
	HIGH_RSTDISBL = 0x80 // Reset Disabled (Enable PC6 as i/o pin)
	HIGH_WDTON    = 0x40 // Watch-dog Timer always on
	HIGH_SPIEN    = 0x20 // Serial program downloading (SPI) enabled
	HIGH_EESAVE   = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ0  = 0x2  // Select Boot Size
	HIGH_BOOTSZ1  = 0x4  // Select Boot Size
	HIGH_BOOTRST  = 0x1  // Boot Reset vector Enabled
	HIGH_CKOPT    = 0x10 // CKOPT fuse (operation dependent of CKSEL fuses)

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

// Bitfields for EXINT: External Interrupts
const (
	// GICR: General Interrupt Control Register
	GICR_INT0  = 0x40 // External Interrupt Request 1 Enable
	GICR_INT1  = 0x80 // External Interrupt Request 1 Enable
	GICR_IVSEL = 0x2  // Interrupt Vector Select
	GICR_IVCE  = 0x1  // Interrupt Vector Change Enable

	// GIFR: General Interrupt Flag Register
	GIFR_INTF0 = 0x40 // External Interrupt Flags
	GIFR_INTF1 = 0x80 // External Interrupt Flags
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR0: Timer/Counter0 Control Register
	TCCR0_CS02 = 0x4 // Clock Select0 bit 2
	TCCR0_CS01 = 0x2 // Clock Select0 bit 1
	TCCR0_CS00 = 0x1 // Clock Select0 bit 0
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A0 = 0x40 // Compare Output Mode 1A, bits
	TCCR1A_COM1A1 = 0x80 // Compare Output Mode 1A, bits
	TCCR1A_COM1B0 = 0x10 // Compare Output Mode 1B, bits
	TCCR1A_COM1B1 = 0x20 // Compare Output Mode 1B, bits
	TCCR1A_FOC1A  = 0x8  // Force Output Compare 1A
	TCCR1A_FOC1B  = 0x4  // Force Output Compare 1B
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
)

// Bitfields for TC8_ASYNC: Timer/Counter, 8-bit Async
const (
	// TCCR2: Timer/Counter2 Control Register
	TCCR2_FOC2  = 0x80 // Force Output Compare
	TCCR2_WGM20 = 0x40 // Waveform Genration Mode
	TCCR2_COM20 = 0x10 // Compare Output Mode bits
	TCCR2_COM21 = 0x20 // Compare Output Mode bits
	TCCR2_WGM21 = 0x8  // Waveform Generation Mode
	TCCR2_CS20  = 0x1  // Clock Select bits
	TCCR2_CS21  = 0x2  // Clock Select bits
	TCCR2_CS22  = 0x4  // Clock Select bits

	// ASSR: Asynchronous Status Register
	ASSR_AS2    = 0x8 // Asynchronous Timer/counter2
	ASSR_TCN2UB = 0x4 // Timer/Counter2 Update Busy
	ASSR_OCR2UB = 0x2 // Output Compare Register2 Update Busy
	ASSR_TCR2UB = 0x1 // Timer/counter Control Register2 Update Busy
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
	UCSRC_URSEL = 0x80 // Register Select
	UCSRC_UMSEL = 0x40 // USART Mode Select
	UCSRC_UPM0  = 0x10 // Parity Mode Bits
	UCSRC_UPM1  = 0x20 // Parity Mode Bits
	UCSRC_USBS  = 0x8  // Stop Bit Select
	UCSRC_UCSZ0 = 0x2  // Character Size
	UCSRC_UCSZ1 = 0x4  // Character Size
	UCSRC_UCPOL = 0x1  // Clock Polarity
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

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCR: Watchdog Timer Control Register
	WDTCR_WDCE = 0x10 // Watchdog Change Enable
	WDTCR_WDE  = 0x8  // Watch Dog Enable
	WDTCR_WDP0 = 0x1  // Watch Dog Timer Prescaler bits
	WDTCR_WDP1 = 0x2  // Watch Dog Timer Prescaler bits
	WDTCR_WDP2 = 0x4  // Watch Dog Timer Prescaler bits
)

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EERIE = 0x8 // EEPROM Ready Interrupt Enable
	EECR_EEMWE = 0x4 // EEPROM Master Write Enable
	EECR_EEWE  = 0x2 // EEPROM Write Enable
	EECR_EERE  = 0x1 // EEPROM Read Enable
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

	// MCUCSR: MCU Control And Status Register
	MCUCSR_WDRF  = 0x8 // Watchdog Reset Flag
	MCUCSR_BORF  = 0x4 // Brown-out Reset Flag
	MCUCSR_EXTRF = 0x2 // External Reset Flag
	MCUCSR_PORF  = 0x1 // Power-on reset flag

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL0 = 0x1  // Oscillator Calibration
	OSCCAL_OSCCAL1 = 0x2  // Oscillator Calibration
	OSCCAL_OSCCAL2 = 0x4  // Oscillator Calibration
	OSCCAL_OSCCAL3 = 0x8  // Oscillator Calibration
	OSCCAL_OSCCAL4 = 0x10 // Oscillator Calibration
	OSCCAL_OSCCAL5 = 0x20 // Oscillator Calibration
	OSCCAL_OSCCAL6 = 0x40 // Oscillator Calibration
	OSCCAL_OSCCAL7 = 0x80 // Oscillator Calibration

	// SPMCR: Store Program Memory Control Register
	SPMCR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCR_RWWSB  = 0x40 // Read-While-Write Section Busy
	SPMCR_RWWSRE = 0x10 // Read-While-Write Section Read Enable
	SPMCR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCR_PGWRT  = 0x4  // Page Write
	SPMCR_PGERS  = 0x2  // Page Erase
	SPMCR_SPMEN  = 0x1  // Store Program Memory Enable
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