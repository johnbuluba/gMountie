<script>

    let tableArr = [{name: "a" }]
    let totalWeight = 0

</script>
<div class="table-container">
    <!-- Native Table Element -->
    <table class="table table-hover">
        <thead>
        <tr>
            <th>Position</th>
            <th>Name</th>
            <th>Symbol</th>
            <th>Weight</th>
        </tr>
        </thead>
        <tbody>
        {#each tableArr as row, i}
            <tr>
                <td>{row.name}</td>
            </tr>
        {/each}
        </tbody>
        <tfoot>
        <tr>
            <th colspan="3">Calculated Total Weight</th>
            <td>{totalWeight}</td>
        </tr>
        </tfoot>
    </table>
</div>
