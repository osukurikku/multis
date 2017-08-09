<template>
  <div class="field has-addons" style="justify-content: flex-end; padding-top: 1.5rem">
    <p class="control">
      <a class="button" :disabled="value <= 0" @click="buttonClick(-1)"><i class="fa fa-arrow-left"></i></a>
    </p>
    <p class="control">
      <a class="button" :disabled="rightDisabled" @click="buttonClick(1)"><i class="fa fa-arrow-right"></i></a>
    </p>
  </div>
</template>

<script>
export default {
  props: ['value', 'rightDisabled'],
  methods: {
    buttonClick (diff) {
      if (diff > 0 && this.rightDisabled) {
        return
      }
      if (diff < 0 && this.value <= 0) {
        return
      }
      this.$emit('input', this.value + diff)
    }
  }
}
</script>

<style>

</style>
