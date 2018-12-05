<template>
    <div>
        <input ref="flp">
    </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import flatpickr from 'flatpickr'
import 'flatpickr/dist/themes/dark.css'

@Component
export default class ExpensesSeeker extends Vue {
  mounted () {
    let element = this.$refs.flp as HTMLElement
    flatpickr(element, {})
  }
}
</script>

<style scoped>

</style>
