<template>
  <div class="p-2">
    <div class="mb-2">Server Stat</div>
    <div><a-button size="mini">Add Server</a-button></div>
    <div v-if="servers">
      <a-table
        :data="servers"
        page-position="tr"
        :pagination="{
          pageSize: pageSize,
          hideOnSinglePage: true,
          size: 'mini',
        }"
        size="small"
      >
        <template #columns>
          <a-table-column title="ID" data-index="id"></a-table-column>
          <a-table-column
            title="Server Name"
            data-index="name"
          ></a-table-column>
          <a-table-column title="URL Server" data-index="url"></a-table-column>
          <a-table-column title="Status Active" align="center">
            <template #cell="{ record }: { record: manga.Server }">
              <span class="chip bg-green-700" v-if="record.status_active"
                >Active</span
              >
              <span v-else class="chip_off">No</span>
            </template>
          </a-table-column>
          <a-table-column title="Chap JSCODE" align="center">
            <template #cell="{ record }: { record: manga.Server }">
              <span
                class="chip on"
                v-if="record.chap_jscode != '' || record.id == 1"
                >SET</span
              >
              <span v-else class="chip off">Not</span>
            </template>
          </a-table-column>
          <a-table-column title="Pages JSCODE" align="center">
            <template #cell="{ record }: { record: manga.Server }">
              <span
                class="chip on"
                v-if="record.page_jscode != '' || record.id == 1"
                >SET</span
              >
              <span v-else class="chip off">Not</span>
            </template>
          </a-table-column>
          <a-table-column title="CRUD">
            <template #cell="{ record }: { record: manga.Server }">
              <a-space>
                <a-button class="h-auto" size="mini">Edit</a-button>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetServer } from '@wails/go/manga/Manga'
import { manga } from '@wails/go/models'
import { GetBreakPoints } from '@/composable/helper'

const servers = ref<manga.Server[] | null>(null)
const fetchServer = async () => {
  servers.value = await GetServer()
}
fetchServer()

const pageSize = ref(16)
const { breakpoints } = GetBreakPoints()
const lg = breakpoints.greater('lg')

if (lg.value) {
  pageSize.value = 23
} else {
  pageSize.value = 16
}

watch(lg, l => {
  if (l) {
    pageSize.value = 23
  } else {
    pageSize.value = 16
  }
})
</script>

<style lang="less" scoped>
.chip {
  border-radius: 0.5rem;
  padding: 0 0.5rem;
  font-size: 0.75rem;
  user-select: none;
}

.on {
  .chip();
  background-color: rgb(24, 136, 101);
}
.off {
  .chip();
  background-color: rgb(167, 16, 16);
}
</style>
