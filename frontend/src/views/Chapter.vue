<template>
  <div>
    Chapter Page
    <br />
    {{ mid }}
  </div>
</template>

<script setup lang="ts">
defineProps<{
  mid: string
}>()
</script>

<style scoped></style>
