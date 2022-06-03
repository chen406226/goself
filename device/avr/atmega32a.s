; Automatically generated file. DO NOT EDIT.
; Generated by gen-device-avr.go from ATmega32A.atdf, see http://packs.download.atmel.com/

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
    jmp __vector_RESET
    jmp __vector_INT0
    jmp __vector_INT1
    jmp __vector_INT2
    jmp __vector_TIMER2_COMP
    jmp __vector_TIMER2_OVF
    jmp __vector_TIMER1_CAPT
    jmp __vector_TIMER1_COMPA
    jmp __vector_TIMER1_COMPB
    jmp __vector_TIMER1_OVF
    jmp __vector_TIMER0_COMP
    jmp __vector_TIMER0_OVF
    jmp __vector_SPI_STC
    jmp __vector_USART_RXC
    jmp __vector_USART_UDRE
    jmp __vector_USART_TXC
    jmp __vector_ADC
    jmp __vector_EE_RDY
    jmp __vector_ANA_COMP
    jmp __vector_TWI
    jmp __vector_SPM_RDY

    ; Define default implementations for interrupts, redirecting to
    ; __vector_default when not implemented.
    IRQ __vector_RESET
    IRQ __vector_INT0
    IRQ __vector_INT1
    IRQ __vector_INT2
    IRQ __vector_TIMER2_COMP
    IRQ __vector_TIMER2_OVF
    IRQ __vector_TIMER1_CAPT
    IRQ __vector_TIMER1_COMPA
    IRQ __vector_TIMER1_COMPB
    IRQ __vector_TIMER1_OVF
    IRQ __vector_TIMER0_COMP
    IRQ __vector_TIMER0_OVF
    IRQ __vector_SPI_STC
    IRQ __vector_USART_RXC
    IRQ __vector_USART_UDRE
    IRQ __vector_USART_TXC
    IRQ __vector_ADC
    IRQ __vector_EE_RDY
    IRQ __vector_ANA_COMP
    IRQ __vector_TWI
    IRQ __vector_SPM_RDY
