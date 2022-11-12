<template>
  <div>
    <img
      @dblclick="$router.push('/test')"
      id="logo"
      alt="Wails logo"
      src="../assets/images/logo-universal.png"
    />
    <HelloWorld />
  </div>
</template>

<script setup lang="ts"></script>

<style scoped></style>
