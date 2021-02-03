<template>
  <table class="shadow-md rounded-md overflow-hidden w-full">
    <tr class="bg-gray-300">
      <slot name="header" />
    </tr>
    <slot name="items" />
    <tr>
      <slot name="footer" />
    </tr>
  </table>
</template>
