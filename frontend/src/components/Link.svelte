<script>
    export let to;
    export let clazz;
</script>

<style>
</style>

<a href={to} class={clazz}>
    <slot/>
</a>