<template>
  <div>
    リダイレクト
  </div>
</template>

<script>
export default {
  validate ({ query, redirect }) {
    // token
    if (query.token === undefined) {
      redirect('/signin')
    } else {
      localStorage.setItem('token', query.token)
    }
    // continue
    if (query.continue === undefined) {
      redirect('/')
    } else {
      redirect(query.continue)
    }
    return true
  }
}
</script>
