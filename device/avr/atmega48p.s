; Automatically generated file. DO NOT EDIT.
; Generated by gen-device-avr.go from ATmega48P.atdf, see http://packs.download.atmel.com/

; This is the default handler for interrupts, if triggered but not defined.
; Sleep inside so that an accidentally triggered interrupt won't drain the
; battery of a battery-powered device.
.section .text.__vector_default
.global  __vector_default
__vector_default:
    sleep
    rjmp __vector_default

; Avoid the need for repeated .weak and .set instructions.
.macro IRQ handler
    .weak  \handler
    .set   \handler, __vector_default
.endm

; The interrupt vector of this device. Must be placed at address 0 by the linker.
.section .vectors
.global  __vectors
    rjmp __vector_RESET
    rjmp __vector_INT0
    rjmp __vector_INT1
    rjmp __vector_PCINT0
    rjmp __vector_PCINT1
    rjmp __vector_PCINT2
    rjmp __vector_WDT
    rjmp __vector_TIMER2_COMPA
    rjmp __vector_TIMER2_COMPB
    rjmp __vector_TIMER2_OVF
    rjmp __vector_TIMER1_CAPT
    rjmp __vector_TIMER1_COMPA
    rjmp __vector_TIMER1_COMPB
    rjmp __vector_TIMER1_OVF
    rjmp __vector_TIMER0_COMPA
    rjmp __vector_TIMER0_COMPB
    rjmp __vector_TIMER0_OVF
    rjmp __vector_SPI_STC
    rjmp __vector_USART_RX
    rjmp __vector_USART_UDRE
    rjmp __vector_USART_TX
    rjmp __vector_ADC
    rjmp __vector_EE_READY
    rjmp __vector_ANALOG_COMP
    rjmp __vector_TWI
    rjmp __vector_SPM_Ready

    ; Define default implementations for interrupts, redirecting to
    ; __vector_default when not implemented.
    IRQ __vector_RESET
    IRQ __vector_INT0
    IRQ __vector_INT1
    IRQ __vector_PCINT0
    IRQ __vector_PCINT1
    IRQ __vector_PCINT2
    IRQ __vector_WDT
    IRQ __vector_TIMER2_COMPA
    IRQ __vector_TIMER2_COMPB
    IRQ __vector_TIMER2_OVF
    IRQ __vector_TIMER1_CAPT
    IRQ __vector_TIMER1_COMPA
    IRQ __vector_TIMER1_COMPB
    IRQ __vector_TIMER1_OVF
    IRQ __vector_TIMER0_COMPA
    IRQ __vector_TIMER0_COMPB
    IRQ __vector_TIMER0_OVF
    IRQ __vector_SPI_STC
    IRQ __vector_USART_RX
    IRQ __vector_USART_UDRE
    IRQ __vector_USART_TX
    IRQ __vector_ADC
    IRQ __vector_EE_READY
    IRQ __vector_ANALOG_COMP
    IRQ __vector_TWI
    IRQ __vector_SPM_Ready
