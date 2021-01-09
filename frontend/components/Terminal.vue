<template>
  <div class="console px-2 py-2">
    <b-row class="mb-2">
      <div class="circle" style="background: var(--orange);" />
      <div class="circle ml-1" style="background: var(--gray);" />
      <div class="circle ml-1" style="background: var(--gray);" />
    </b-row>

    <div>
      $
      <div class="typing" style="display:inline-block; vertical-align: top;">
        {{ input }}
      </div>
    </div>
    <div v-show="enterd">
      {{ response }}
    </div>
  </div>
</template>

<script>
export default {
  props: {
    command: {
      type: String,
      default: 'shub search \'ping\''
    },
    response: {
      type: String,
      default: 'pong'
    },
    speed: {
      type: Number,
      default: 150
    }
  },
  data () {
    return {
      input: '',
      enterd: false
    }
  },
  created () {
    // typewriter
    for (let i = 0; i < this.command.length; i++) {
      setTimeout(() => {
        this.input += this.command[i]
      }, this.speed * i)
    }
    // result
    setTimeout(() => {
      this.enterd = true
    }, this.speed * this.command.length)
  }
}
</script>

<style lang="scss" scoped>
.console {
  font-family: Consolas;
  font-size: 14px;
  color: $white;
  background: $navyblue;
  width: 100%;
  height: 360px;
  border-radius: 10px;
}

@keyframes cursor {
  50% { border-right-color: transparent; }
}

.typing {
  width: auto;
  border-right: 1ch solid;
  animation: cursor 1s steps(1) infinite;
}

.circle {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}
</style>
