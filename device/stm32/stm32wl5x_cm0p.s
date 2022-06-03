// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from stm32wl5x_cm0p.svd, see https://github.com/tinygo-org/stm32-svd

// STM32WL5x_CM0P
//


.syntax unified

// This is the default handler for interrupts, if triggered but not defined.
.section .text.Default_Handler
.global  Default_Handler
.type    Default_Handler, %function
Default_Handler:
    wfe
    b    Default_Handler
.size Default_Handler, .-Default_Handler

// Avoid the need for repeated .weak and .set instructions.
.macro IRQ handler
    .weak  \handler
    .set   \handler, Default_Handler
.endm

// Must set the "a" flag on the section:
// https://svnweb.freebsd.org/base/stable/11/sys/arm/arm/locore-v4.S?r1=321049&r2=321048&pathrev=321049
// https://sourceware.org/binutils/docs/as/Section.html#ELF-Version
.section .isr_vector, "a", %progbits
.global  __isr_vector
__isr_vector:
    // Interrupt vector as defined by Cortex-M, starting with the stack top.
    // On reset, SP is initialized with *0x0 and PC is loaded with *0x4, loading
    // _stack_top and Reset_Handler.
    .long _stack_top
    .long Reset_Handler
    .long NMI_Handler
    .long HardFault_Handler
    .long MemoryManagement_Handler
    .long BusFault_Handler
    .long UsageFault_Handler
    .long 0
    .long 0
    .long 0
    .long 0
    .long SVC_Handler
    .long DebugMon_Handler
    .long 0
    .long PendSV_Handler
    .long SysTick_Handler

    // Extra interrupts for peripherals defined by the hardware vendor.
    .long TZIC_ILA_IRQHandler
    .long PVD_PVM_3_IRQHandler
    .long TAMP_RTCSTAMP_LSECSS_RTCALARM_RTCSSRU_RTCWKUP_IRQHandler
    .long FLASH_RCC_C1SEV_IRQHandler
    .long EXTI1_0_IRQHandler
    .long EXTI3_2_IRQHandler
    .long EXTI15_4_IRQHandler
    .long ADC_COMP_DAC_IRQHandler
    .long DMA1_CH3_1_IRQHandler
    .long DMA1_CH7_4_IRQHandler
    .long DMA2_CH7_1_DMAMUX1_OVR_IRQHandler
    .long LPTIM1_IRQHandler
    .long LPTIM2_IRQHandler
    .long LPTIM3_IRQHandler
    .long TIM1_BRK_TIM1_UP_TIM1_TRG_COM_TIM1_CC_IRQHandler
    .long TIM2_IRQHandler
    .long TIM16_IRQHandler
    .long TIM17_IRQHandler
    .long IPCC_C2_RX_IT_IPCC_C2_TX_IT_IRQHandler
    .long HSEM_IRQHandler
    .long True_RNG_IRQHandler
    .long AES_PKA_IRQHandler
    .long I2C1_EV_I2C1_ER_IRQHandler
    .long I2C2_EV_I2C2_ER_IRQHandler
    .long I2C3_EV_I2C3_ER_IRQHandler
    .long SPI1_IRQHandler
    .long SPI2S2_IRQHandler
    .long USART1_IRQHandler
    .long USART2_IRQHandler
    .long LPUART1_IRQHandler
    .long SUBGHZSPI_IRQHandler
    .long Radio_IRQ_Busy_IRQHandler

    // Define default implementations for interrupts, redirecting to
    // Default_Handler when not implemented.
    IRQ NMI_Handler
    IRQ HardFault_Handler
    IRQ MemoryManagement_Handler
    IRQ BusFault_Handler
    IRQ UsageFault_Handler
    IRQ SVC_Handler
    IRQ DebugMon_Handler
    IRQ PendSV_Handler
    IRQ SysTick_Handler
    IRQ TZIC_ILA_IRQHandler
    IRQ PVD_PVM_3_IRQHandler
    IRQ TAMP_RTCSTAMP_LSECSS_RTCALARM_RTCSSRU_RTCWKUP_IRQHandler
    IRQ FLASH_RCC_C1SEV_IRQHandler
    IRQ EXTI1_0_IRQHandler
    IRQ EXTI3_2_IRQHandler
    IRQ EXTI15_4_IRQHandler
    IRQ ADC_COMP_DAC_IRQHandler
    IRQ DMA1_CH3_1_IRQHandler
    IRQ DMA1_CH7_4_IRQHandler
    IRQ DMA2_CH7_1_DMAMUX1_OVR_IRQHandler
    IRQ LPTIM1_IRQHandler
    IRQ LPTIM2_IRQHandler
    IRQ LPTIM3_IRQHandler
    IRQ TIM1_BRK_TIM1_UP_TIM1_TRG_COM_TIM1_CC_IRQHandler
    IRQ TIM2_IRQHandler
    IRQ TIM16_IRQHandler
    IRQ TIM17_IRQHandler
    IRQ IPCC_C2_RX_IT_IPCC_C2_TX_IT_IRQHandler
    IRQ HSEM_IRQHandler
    IRQ True_RNG_IRQHandler
    IRQ AES_PKA_IRQHandler
    IRQ I2C1_EV_I2C1_ER_IRQHandler
    IRQ I2C2_EV_I2C2_ER_IRQHandler
    IRQ I2C3_EV_I2C3_ER_IRQHandler
    IRQ SPI1_IRQHandler
    IRQ SPI2S2_IRQHandler
    IRQ USART1_IRQHandler
    IRQ USART2_IRQHandler
    IRQ LPUART1_IRQHandler
    IRQ SUBGHZSPI_IRQHandler
    IRQ Radio_IRQ_Busy_IRQHandler
