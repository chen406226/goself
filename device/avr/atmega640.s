; Automatically generated file. DO NOT EDIT.
; Generated by gen-device-avr.go from ATmega640.atdf, see http://packs.download.atmel.com/

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
    jmp __vector_INT3
    jmp __vector_INT4
    jmp __vector_INT5
    jmp __vector_INT6
    jmp __vector_INT7
    jmp __vector_PCINT0
    jmp __vector_PCINT1
    jmp __vector_PCINT2
    jmp __vector_WDT
    jmp __vector_TIMER2_COMPA
    jmp __vector_TIMER2_COMPB
    jmp __vector_TIMER2_OVF
    jmp __vector_TIMER1_CAPT
    jmp __vector_TIMER1_COMPA
    jmp __vector_TIMER1_COMPB
    jmp __vector_TIMER1_COMPC
    jmp __vector_TIMER1_OVF
    jmp __vector_TIMER0_COMPA
    jmp __vector_TIMER0_COMPB
    jmp __vector_TIMER0_OVF
    jmp __vector_SPI_STC
    jmp __vector_USART0_RX
    jmp __vector_USART0_UDRE
    jmp __vector_USART0_TX
    jmp __vector_ANALOG_COMP
    jmp __vector_ADC
    jmp __vector_EE_READY
    jmp __vector_TIMER3_CAPT
    jmp __vector_TIMER3_COMPA
    jmp __vector_TIMER3_COMPB
    jmp __vector_TIMER3_COMPC
    jmp __vector_TIMER3_OVF
    jmp __vector_USART1_RX
    jmp __vector_USART1_UDRE
    jmp __vector_USART1_TX
    jmp __vector_TWI
    jmp __vector_SPM_READY
    jmp __vector_TIMER4_CAPT
    jmp __vector_TIMER4_COMPA
    jmp __vector_TIMER4_COMPB
    jmp __vector_TIMER4_COMPC
    jmp __vector_TIMER4_OVF
    jmp __vector_TIMER5_CAPT
    jmp __vector_TIMER5_COMPA
    jmp __vector_TIMER5_COMPB
    jmp __vector_TIMER5_COMPC
    jmp __vector_TIMER5_OVF
    jmp __vector_USART2_RX
    jmp __vector_USART2_UDRE
    jmp __vector_USART2_TX
    jmp __vector_USART3_RX
    jmp __vector_USART3_UDRE
    jmp __vector_USART3_TX

    ; Define default implementations for interrupts, redirecting to
    ; __vector_default when not implemented.
    IRQ __vector_RESET
    IRQ __vector_INT0
    IRQ __vector_INT1
    IRQ __vector_INT2
    IRQ __vector_INT3
    IRQ __vector_INT4
    IRQ __vector_INT5
    IRQ __vector_INT6
    IRQ __vector_INT7
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
    IRQ __vector_TIMER1_COMPC
    IRQ __vector_TIMER1_OVF
    IRQ __vector_TIMER0_COMPA
    IRQ __vector_TIMER0_COMPB
    IRQ __vector_TIMER0_OVF
    IRQ __vector_SPI_STC
    IRQ __vector_USART0_RX
    IRQ __vector_USART0_UDRE
    IRQ __vector_USART0_TX
    IRQ __vector_ANALOG_COMP
    IRQ __vector_ADC
    IRQ __vector_EE_READY
    IRQ __vector_TIMER3_CAPT
    IRQ __vector_TIMER3_COMPA
    IRQ __vector_TIMER3_COMPB
    IRQ __vector_TIMER3_COMPC
    IRQ __vector_TIMER3_OVF
    IRQ __vector_USART1_RX
    IRQ __vector_USART1_UDRE
    IRQ __vector_USART1_TX
    IRQ __vector_TWI
    IRQ __vector_SPM_READY
    IRQ __vector_TIMER4_CAPT
    IRQ __vector_TIMER4_COMPA
    IRQ __vector_TIMER4_COMPB
    IRQ __vector_TIMER4_COMPC
    IRQ __vector_TIMER4_OVF
    IRQ __vector_TIMER5_CAPT
    IRQ __vector_TIMER5_COMPA
    IRQ __vector_TIMER5_COMPB
    IRQ __vector_TIMER5_COMPC
    IRQ __vector_TIMER5_OVF
    IRQ __vector_USART2_RX
    IRQ __vector_USART2_UDRE
    IRQ __vector_USART2_TX
    IRQ __vector_USART3_RX
    IRQ __vector_USART3_UDRE
    IRQ __vector_USART3_TX
