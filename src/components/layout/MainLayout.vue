<q-tabs
  v-model="activeTab"
  class="bg-primary text-white shadow-2"
  active-color="white"
  indicator-color="white"
  align="justify"
>
  <q-route-tab
    name="dashboard"
    icon="dashboard"
    :to="{ name: 'dashboard' }"
    label="首页"
  />
  <q-route-tab
    name="community"
    icon="community"
    :to="{ name: 'community' }"
    label="社区"
  />
  <q-route-tab
    name="profile"
    icon="profile"
    :to="{ name: 'profile' }"
    label="我的"
  />
</q-tabs>
