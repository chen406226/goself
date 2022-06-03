; Automatically generated file. DO NOT EDIT.
; Generated by gen-device-avr.go from ATtiny44A.atdf, see http://packs.download.atmel.com/

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
    rjmp __vector_EXT_INT0
    rjmp __vector_PCINT0
    rjmp __vector_PCINT1
    rjmp __vector_WDT
    rjmp __vector_TIM1_CAPT
    rjmp __vector_TIM1_COMPA
    rjmp __vector_TIM1_COMPB
    rjmp __vector_TIM1_OVF
    rjmp __vector_TIM0_COMPA
    rjmp __vector_TIM0_COMPB
    rjmp __vector_TIM0_OVF
    rjmp __vector_ANA_COMP
    rjmp __vector_ADC
    rjmp __vector_EE_RDY
    rjmp __vector_USI_STR
    rjmp __vector_USI_OVF

    ; Define default implementations for interrupts, redirecting to
    ; __vector_default when not implemented.
    IRQ __vector_RESET
    IRQ __vector_EXT_INT0
    IRQ __vector_PCINT0
    IRQ __vector_PCINT1
    IRQ __vector_WDT
    IRQ __vector_TIM1_CAPT
    IRQ __vector_TIM1_COMPA
    IRQ __vector_TIM1_COMPB
    IRQ __vector_TIM1_OVF
    IRQ __vector_TIM0_COMPA
    IRQ __vector_TIM0_COMPB
    IRQ __vector_TIM0_OVF
    IRQ __vector_ANA_COMP
    IRQ __vector_ADC
    IRQ __vector_EE_RDY
    IRQ __vector_USI_STR
    IRQ __vector_USI_OVF
