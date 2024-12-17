<template>
  <div class="bg-white">
    <!-- Mobile menu -->
    <TransitionRoot as="template" :show="open">
      <Dialog class="relative z-40 lg:hidden" @close="open = false">
        <TransitionChild as="template" enter="transition-opacity ease-linear duration-300" enter-from="opacity-0"
          enter-to="opacity-100" leave="transition-opacity ease-linear duration-300" leave-from="opacity-100"
          leave-to="opacity-0">
          <div class="fixed inset-0 bg-black/25" />
        </TransitionChild>

        <div class="fixed inset-0 z-40 flex">
          <TransitionChild as="template" enter="transition ease-in-out duration-300 transform"
            enter-from="-translate-x-full" enter-to="translate-x-0"
            leave="transition ease-in-out duration-300 transform" leave-from="translate-x-0"
            leave-to="-translate-x-full">
            <DialogPanel class="relative flex w-full max-w-xs flex-col overflow-y-auto bg-white pb-12 shadow-xl">
              <div class="flex px-4 pb-2 pt-5">
                <button type="button"
                  class="relative -m-2 inline-flex items-center justify-center rounded-md p-2 text-gray-400"
                  @click="open = false">
                  <span class="absolute -inset-0.5" />
                  <span class="sr-only">Close menu</span>
                  <XMarkIcon class="size-6" aria-hidden="true" />
                </button>
              </div>

              <!-- Links -->
              <TabGroup as="div" class="mt-2">
                <div class="border-b border-gray-200">
                  <TabList class="-mb-px flex space-x-8 px-4">
                    <Tab as="template" v-for="category in navigation.categories" :key="category.name"
                      v-slot="{ selected }">
                      <button
                        :class="[selected ? 'border-indigo-600 text-indigo-600' : 'border-transparent text-gray-900', 'flex-1 whitespace-nowrap border-b-2 px-1 py-4 text-base font-medium']">{{
                          category.name }}</button>
                    </Tab>
                  </TabList>
                </div>
                <TabPanels as="template">
                  <TabPanel v-for="category in navigation.categories" :key="category.name"
                    class="space-y-10 px-4 pb-8 pt-10">
                    <div class="grid grid-cols-2 gap-x-4">
                      <div v-for="item in category.featured" :key="item.name" class="group relative text-sm">
                        <img :src="item.imageSrc" :alt="item.imageAlt"
                          class="aspect-square w-full rounded-lg bg-gray-100 object-cover group-hover:opacity-75" />
                        <a :href="item.href" class="mt-6 block font-medium text-gray-900">
                          <span class="absolute inset-0 z-10" aria-hidden="true" />
                          {{ item.name }}
                        </a>
                        <p aria-hidden="true" class="mt-1">Shop now</p>
                      </div>
                    </div>
                    <div v-for="section in category.sections" :key="section.name">
                      <p :id="`${category.id}-${section.id}-heading-mobile`" class="font-medium text-gray-900">{{
                        section.name }}</p>
                      <ul role="list" :aria-labelledby="`${category.id}-${section.id}-heading-mobile`"
                        class="mt-6 flex flex-col space-y-6">
                        <li v-for="item in section.items" :key="item.name" class="flow-root">
                          <a :href="item.href" class="-m-2 block p-2 text-gray-500">{{ item.name }}</a>
                        </li>
                      </ul>
                    </div>
                  </TabPanel>
                </TabPanels>
              </TabGroup>

              <div class="space-y-6 border-t border-gray-200 px-4 py-6">
                <div v-for="page in navigation.pages" :key="page.name" class="flow-root">
                  <a :href="page.href" class="-m-2 block p-2 font-medium text-gray-900">{{ page.name }}</a>
                </div>
              </div>

              <div class="space-y-6 border-t border-gray-200 px-4 py-6">
                <div class="flow-root">
                  <a href="#" class="-m-2 block p-2 font-medium text-gray-900">Sign in</a>
                </div>
                <div class="flow-root">
                  <a href="#" class="-m-2 block p-2 font-medium text-gray-900">Create account</a>
                </div>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </Dialog>
    </TransitionRoot>

    <header class=" relative bg-white">
      <p class="flex h-10 items-center justify-center bg-rose-600 px-4 text-sm font-medium text-white sm:px-6 lg:px-8">
        Miễn phí vận chuyển cho đơn hàng trên 100.000đ</p>

      <nav aria-label="Top" class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div class="border-b border-gray-200">
          <div class="flex h-24 items-center">
            <button type="button" class="relative rounded-md bg-white p-2 text-gray-400 lg:hidden" @click="open = true">
              <span class="absolute -inset-0.5" />
              <span class="sr-only">Open menu</span>
              <Bars3Icon class="size-6" aria-hidden="true" />
            </button>

            <!-- Logo -->
            <div class="ml-4 flex lg:ml-0">
              <span class="sr-only">Your Company</span>
              <RouterLink :to="{ name: 'home' }" class="active">
                <img class="h-20 w-auto" src="../assets/image/Logo.png" alt="" />
              </RouterLink>
            </div>

            <!-- Flyout menus -->
            <PopoverGroup class="hidden lg:ml-8 lg:block lg:self-stretch">
              <div class="flex h-full space-x-8">
                <Popover v-for="category in navigation.categories" :key="category.name" class="flex" v-slot="{ open }">
                  <div class="relative flex">
                    <PopoverButton
                      :class="[open ? 'border-red-500 text-rose-600' : 'border-transparent text-gray-700 hover:text-gray-800', 'relative z-10 -mb-px flex items-center border-b-2 pt-px text-sm font-medium transition-colors duration-200 ease-out']">
                      {{ category.name }}</PopoverButton>
                  </div>

                  <transition enter-active-class="transition ease-out duration-200" enter-from-class="opacity-0"
                    enter-to-class="opacity-100" leave-active-class="transition ease-in duration-150"
                    leave-from-class="opacity-100" leave-to-class="opacity-0">
                    <PopoverPanel class="absolute inset-x-0 top-full text-sm text-gray-500">
                      <!-- Presentational element used to render the bottom shadow, if we put the shadow on the actual panel it pokes out the top, so we use this shorter element to hide the top of the shadow -->
                      <div class="absolute inset-0 top-1/2 bg-white shadow" aria-hidden="true" />

                      <div class="relative bg-white">
                        <div class="mx-auto max-w-7xl px-8">
                          <div class="grid grid-cols-2 gap-x-8 gap-y-10 py-16">
                            <div class="col-start-2 grid grid-cols-2 gap-x-8">
                              <div v-for="item in category.featured" :key="item.name"
                                class="group relative text-base sm:text-sm">
                                <img :src="item.imageSrc" :alt="item.imageAlt"
                                  class="aspect-square w-full rounded-lg bg-gray-100 object-cover group-hover:opacity-75" />
                                <a :href="item.href" class="mt-6 block font-medium text-gray-900">
                                  <span class="absolute inset-0 z-10" aria-hidden="true" />
                                  {{ item.name }}
                                </a>
                                <p aria-hidden="true" class="mt-1">Shop now</p>
                              </div>
                            </div>
                            <div class="row-start-1 grid grid-cols-3 gap-x-8 gap-y-10 text-sm">
                              <div v-for="section in category.sections" :key="section.name">
                                <p :id="`${section.name}-heading`" class="font-medium text-gray-900">{{ section.name }}
                                </p>
                                <ul role="list" :aria-labelledby="`${section.name}-heading`"
                                  class="mt-6 space-y-6 sm:mt-4 sm:space-y-4">
                                  <li v-for="item in section.items" :key="item.name" class="flex">
                                    <a :href="item.href" class="hover:text-gray-800">{{ item.name }}</a>
                                  </li>
                                </ul>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </PopoverPanel>
                  </transition>
                </Popover>

                <a v-for="page in navigation.pages" :key="page.name" :href="page.href"
                  class="flex items-center text-sm font-medium text-gray-700 hover:text-gray-800">{{ page.name }}</a>
              </div>
            </PopoverGroup>

            <div class="ml-auto flex items-center">
              <div class="hidden lg:flex lg:flex-1 lg:items-center lg:justify-end lg:space-x-6 mr-2">
                <router-link :to="{ name: 'signIn' }" class="text-sm font-medium text-gray-700 hover:text-gray-800">Đăng
                  nhập</router-link>
                <span class="h-6 w-px bg-gray-200" aria-hidden="true" />
                <router-link :to="{ name: 'signUp' }" class="text-sm font-medium text-gray-700 hover:text-gray-800">Đăng
                  ký</router-link>
              </div>


              <!-- Search -->
              <div>
                <div class="relative text-gray-600">
                  <input type="search" name="serch" placeholder="Search"
                    class="bg-white h-10 px-5 pr-10 rounded-full text-sm focus:outline-none">
                  <button type="submit" class="absolute right-0 top-0 mt-3 mr-4">
                    <svg class="h-4 w-4 fill-current" xmlns="http://www.w3.org/2000/svg"
                      xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" id="Capa_1" x="0px" y="0px"
                      viewBox="0 0 56.966 56.966" style="enable-background:new 0 0 56.966 56.966;" xml:space="preserve"
                      width="512px" height="512px">
                      <path
                        d="M55.146,51.887L41.588,37.786c3.486-4.144,5.396-9.358,5.396-14.786c0-12.682-10.318-23-23-23s-23,10.318-23,23  s10.318,23,23,23c4.761,0,9.298-1.436,13.177-4.162l13.661,14.208c0.571,0.593,1.339,0.92,2.162,0.92  c0.779,0,1.518-0.297,2.079-0.837C56.255,54.982,56.293,53.08,55.146,51.887z M23.984,6c9.374,0,17,7.626,17,17s-7.626,17-17,17  s-17-7.626-17-17S14.61,6,23.984,6z" />
                    </svg>
                  </button>
                </div>
              </div>

              <!-- Cart -->
              <div class="ml-4 flow-root lg:ml-6">
                <a href="#" class="group -m-2 flex items-center p-2">
                  <ShoppingBagIcon class="size-6 shrink-0 text-gray-400 group-hover:text-gray-500" aria-hidden="true" />
                  <span class="ml-2 text-sm font-medium text-gray-700 group-hover:text-gray-800">0</span>
                  <span class="sr-only">items in cart, view bag</span>
                </a>
              </div>
            </div>
          </div>
        </div>
      </nav>
    </header>
  </div>
</template>

<script>
import {
  Dialog,
  DialogPanel,
  Popover,
  PopoverButton,
  PopoverGroup,
  PopoverPanel,
  Tab,
  TabGroup,
  TabList,
  TabPanel,
  TabPanels,
  TransitionChild,
  TransitionRoot,
} from '@headlessui/vue';
import { Bars3Icon, MagnifyingGlassIcon, ShoppingBagIcon, XMarkIcon } from '@heroicons/vue/24/outline'

export default {
  name: 'Navigation',
  components: {
    TransitionRoot,
    TransitionChild,
    Dialog,
    DialogPanel,
    TabGroup,
    TabList,
    Tab,
    XMarkIcon,
  },
  data() {
    return {
      open: false,
      navigation: {
        categories: [
          {
            id: 'women',
            name: 'Nữ',
            featured: [
              {
                name: 'Sản phẩm mới',
                href: '#',
                imageSrc: 'https://mir-s3-cdn-cf.behance.net/project_modules/max_1200/936b4f212945879.673ea613191e9.jpg',
                imageAlt: 'Models sitting back to back, wearing Basic Tee in black and bone.',
              },
              {
                name: 'Lego',
                href: '#',
                imageSrc: 'https://www.behance.net/gallery/213115183/Sensory-and-Virtual-Design-LEGo',
                imageAlt: 'Close up of Basic Tee fall bundle with off-white, ochre, olive, and black tees.',
              },
            ],
            sections: [
              {
                id: 'toy',
                name: 'Đồ chơi',
                items: [
                  { name: 'Tops', href: '#' },
                  { name: 'Dresses', href: '#' },
                  { name: 'Pants', href: '#' },

                ],
              },
              {
                id: 'doll',
                name: 'Búp bê',
                items: [
                  { name: 'Búp bê thời trang', href: '#' },
                  { name: 'Búp bê baby', href: '#' },
                  { name: 'Búp bê sưu tập', href: '#' },
                ],
              },
              {
                id: 'brands',
                name: 'Thương hiệu',
                items: [
                  { name: 'Lego', href: '#' },
                  { name: 'AngryBird', href: '#' },
                  { name: 'Barbie', href: '#' },
                  { name: 'Iwako', href: '#' },
                  { name: 'Hatchimals', href: '#' },
                ],
              },
            ],
          },
          {
            id: 'men',
            name: 'Nam',
            featured: [
              {
                name: 'Sản phẩm mới',
                href: '#',
                imageSrc: '',
                imageAlt: 'Drawstring top with elastic loop closure and textured interior padding.',
              },
              {
                name: 'Lego',
                href: '#',
                imageSrc: '',
                imageAlt:
                  'Three shirts in gray, white, and blue arranged on table with same line drawing of hands and shapes overlapping on front of shirt.',
              },
            ],
            sections: [
              {
                id: 'clothing',
                name: 'Clothing',
                items: [
                  { name: 'Tops', href: '#' },
                  { name: 'Pants', href: '#' },
                  { name: 'Sweaters', href: '#' },

                ],
              },
              {
                id: 'accessories',
                name: 'Accessories',
                items: [
                  { name: 'Watches', href: '#' },
                  { name: 'Wallets', href: '#' },
                  { name: 'Bags', href: '#' },

                ],
              },
              {
                id: 'brands',
                name: 'Brands',
                items: [
                  { name: 'Re-Arranged', href: '#' },
                  { name: 'Counterfeit', href: '#' },
                  { name: 'Full Nelson', href: '#' },
                  { name: 'My Way', href: '#' },
                ],
              },
            ],
          },
        ],
        pages: [
          { name: 'Thương hiệu', href: '#' },
          { name: 'Hàng mới', href: '#' },
        ],
      }
    }
  }
}


</script>