<template>
  <div id="clock">
    <p>Текущее время: {{now.toLocaleTimeString('ru-RU')}}</p>
    <p>Актуальность: {{new Date(ts).toLocaleTimeString('ru-RU')}}</p>
  </div>
</template>

<script>
export default {
  name: 'clock',
  props: {
    timestamp: Date,
  },
  data(){
    return {
      now: new Date(),
      ts: this.timestamp,
    }
  }
}
</script>

<style scoped>
#clock {
  top: 1px;
  right: 1px;
  position: absolute;
}
</style>
