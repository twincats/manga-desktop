<template>
  <th>
    <div v-if="column.title == 'Download'" class="text-center">
      <a-button
        @click="emit('clickDownload')"
        type="primary"
        size="mini"
        v-if="active"
        >Download</a-button
      >
      <slot v-else />
    </div>
    <div v-else>
      <slot />
    </div>
  </th>
</template>

<script setup lang="ts">
defineProps<{
  column: any
  active: boolean
}>()

const emit = defineEmits<{
  (e: 'clickDownload'): void
}>()
</script>

<style scoped></style>
