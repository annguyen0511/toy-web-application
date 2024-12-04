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
      <p
        class="flex h-10 items-center justify-center bg-rose-600 px-4 text-sm font-medium text-white sm:px-6 lg:px-8">
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
                <RouterLink :to="{name : 'home'}" class="active">
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
                <router-link :to="{name: 'signIn' }" class="text-sm font-medium text-gray-700 hover:text-gray-800">Đăng nhập</router-link>
                <span class="h-6 w-px bg-gray-200" aria-hidden="true" />
                <router-link :to="{name: 'signUp'}" class="text-sm font-medium text-gray-700 hover:text-gray-800">Đăng ký</router-link>
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
  <div class="bg-white">
    <div class="mx-auto max-w-2xl px-4 py-16 sm:px-6 sm:py-24 lg:max-w-7xl lg:px-8">
      <h2 class="sr-only">Products</h2>

      <div class="grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 xl:gap-x-8">
        <a v-for="product in products" :key="product.id" :href="product.href" class="group">
          <img :src="product.imageSrc" :alt="product.imageAlt" class="aspect-square w-full rounded-lg bg-gray-200 object-cover group-hover:opacity-75 xl:aspect-[7/8]" />
          <h3 class="mt-4 text-sm text-gray-700">{{ product.name }}</h3>
          <p class="mt-1 text-lg font-medium text-gray-900">{{ product.price }}</p>
        </a>
      </div>
      <div class="flex items-center justify-between border-t border-gray-200 bg-white px-4 py-3 sm:px-6">
    <div class="flex flex-1 justify-between sm:hidden">
      <a href="#" class="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">Previous</a>
      <a href="#" class="relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">Next</a>
    </div>
    <div class="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
      <div>
        <p class="text-sm text-gray-700">
          Showing
          {{ ' ' }}
          <span class="font-medium">1</span>
          {{ ' ' }}
          to
          {{ ' ' }}
          <span class="font-medium">10</span>
          {{ ' ' }}
          of
          {{ ' ' }}
          <span class="font-medium">97</span>
          {{ ' ' }}
          results
        </p>
      </div>
      <div>
        <nav class="isolate inline-flex -space-x-px rounded-md shadow-sm" aria-label="Pagination">
          <a href="#" class="relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0">
            <span class="sr-only">Previous</span>
            <ChevronLeftIcon class="size-5" aria-hidden="true" />
          </a>
          <!-- Current: "z-10 bg-indigo-600 text-white focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600", Default: "text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:outline-offset-0" -->
          <a href="#" aria-current="page" class="relative z-10 inline-flex items-center bg-rose-500 px-4 py-2 text-sm font-semibold text-white focus:z-20 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">1</a>
          <a href="#" class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0">2</a>
          <a href="#" class="relative hidden items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0 md:inline-flex">3</a>
          <span class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-700 ring-1 ring-inset ring-gray-300 focus:outline-offset-0">...</span>
          <a href="#" class="relative hidden items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0 md:inline-flex">8</a>
          <a href="#" class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0">9</a>
          <a href="#" class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0">10</a>
          <a href="#" class="relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0">
            <span class="sr-only">Next</span>
            <ChevronRightIcon class="size-5" aria-hidden="true" />
          </a>
        </nav>
      </div>
    </div>
  </div>
    </div>
  </div>
  <div class="bg-white py-24 sm:py-32">
    <div class="mx-auto max-w-7xl px-6 lg:px-8">
      <h2 class="text-center text-5xl font-bold text-rose-500">Thương Hiệu Nổi Bật</h2>
      <div class="mx-auto mt-10 grid max-w-lg grid-cols-4 items-center gap-x-8 gap-y-10 sm:max-w-xl sm:grid-cols-6 sm:gap-x-10 lg:mx-0 lg:max-w-none lg:grid-cols-5">
        <img class="col-span-2 max-h-12 w-full object-contain lg:col-span-1" src="https://tailwindui.com/plus/img/logos/158x48/transistor-logo-gray-900.svg" alt="Transistor" width="158" height="48" />
        <img class="col-span-2 max-h-12 w-full object-contain lg:col-span-1" src="https://tailwindui.com/plus/img/logos/158x48/reform-logo-gray-900.svg" alt="Reform" width="158" height="48" />
        <img class="col-span-2 max-h-12 w-full object-contain lg:col-span-1" src="https://tailwindui.com/plus/img/logos/158x48/tuple-logo-gray-900.svg" alt="Tuple" width="158" height="48" />
        <img class="col-span-2 max-h-12 w-full object-contain sm:col-start-2 lg:col-span-1" src="https://tailwindui.com/plus/img/logos/158x48/savvycal-logo-gray-900.svg" alt="SavvyCal" width="158" height="48" />
        <img class="col-span-2 col-start-2 max-h-12 w-full object-contain sm:col-start-auto lg:col-span-1" src="https://tailwindui.com/plus/img/logos/158x48/statamic-logo-gray-900.svg" alt="Statamic" width="158" height="48" />
      </div>
    </div>
  </div>
  

<footer class="bg-white dark:bg-gray-900">
    <div class="mx-auto w-full max-w-screen-xl p-4 py-6 lg:py-8">
        <div class="md:flex md:justify-between">
          <div class="mb-6 md:mb-0">
              <a href="https://flowbite.com/" class="flex items-center">
                  <img src="../assets/image/Logo.png" class="h-8 me-3" alt="FlowBite Logo" />
                  <span class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Toy Toy World</span>
              </a>
          </div>
          <div class="grid grid-cols-2 gap-8 sm:gap-6 sm:grid-cols-3">
              <div>
              </div>
              <div>
                  <h2 class="mb-6 text-sm font-semibold text-gray-900 uppercase dark:text-white">Theo dõi chúng tôi</h2>
                  <ul class="text-gray-500 dark:text-gray-400 font-medium">
                      <li>
                          <a href="https://discord.gg/h3eFpmaK" class="hover:underline">Discord</a>
                      </li>
                  </ul>
              </div>
              <div>
                  <h2 class="mb-6 text-sm font-semibold text-gray-900 uppercase dark:text-white">Điều khoản</h2>
                  <ul class="text-gray-500 dark:text-gray-400 font-medium">
                      <li class="mb-4">
                          <a href="#" class="hover:underline">Privacy Policy</a>
                      </li>
                      <li>
                          <a href="#" class="hover:underline">Terms &amp; Conditions</a>
                      </li>
                  </ul>
              </div>
          </div>
      </div>
      <hr class="my-6 border-gray-200 sm:mx-auto dark:border-gray-700 lg:my-8" />
      <div class="sm:flex sm:items-center sm:justify-between">
          <span class="text-sm text-gray-500 sm:text-center dark:text-gray-400">© 2024 <a href="" class="hover:underline">Toy Toy World</a>. Đã đăng ký bản quyền.
          </span>
          <div class="flex mt-4 sm:justify-center sm:mt-0">
              <a href="#" class="text-gray-500 hover:text-gray-900 dark:hover:text-white">
                  <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 8 19">
                        <path fill-rule="evenodd" d="M6.135 3H8V0H6.135a4.147 4.147 0 0 0-4.142 4.142V6H0v3h2v9.938h3V9h2.021l.592-3H5V3.591A.6.6 0 0 1 5.592 3h.543Z" clip-rule="evenodd"/>
                    </svg>
                  <span class="sr-only">Facebook page</span>
              </a>
              <a href="#" class="text-gray-500 hover:text-gray-900 dark:hover:text-white ms-5">
                  <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 21 16">
                        <path d="M16.942 1.556a16.3 16.3 0 0 0-4.126-1.3 12.04 12.04 0 0 0-.529 1.1 15.175 15.175 0 0 0-4.573 0 11.585 11.585 0 0 0-.535-1.1 16.274 16.274 0 0 0-4.129 1.3A17.392 17.392 0 0 0 .182 13.218a15.785 15.785 0 0 0 4.963 2.521c.41-.564.773-1.16 1.084-1.785a10.63 10.63 0 0 1-1.706-.83c.143-.106.283-.217.418-.33a11.664 11.664 0 0 0 10.118 0c.137.113.277.224.418.33-.544.328-1.116.606-1.71.832a12.52 12.52 0 0 0 1.084 1.785 16.46 16.46 0 0 0 5.064-2.595 17.286 17.286 0 0 0-2.973-11.59ZM6.678 10.813a1.941 1.941 0 0 1-1.8-2.045 1.93 1.93 0 0 1 1.8-2.047 1.919 1.919 0 0 1 1.8 2.047 1.93 1.93 0 0 1-1.8 2.045Zm6.644 0a1.94 1.94 0 0 1-1.8-2.045 1.93 1.93 0 0 1 1.8-2.047 1.918 1.918 0 0 1 1.8 2.047 1.93 1.93 0 0 1-1.8 2.045Z"/>
                    </svg>
                  <span class="sr-only">Discord community</span>
              </a>
              <a href="#" class="text-gray-500 hover:text-gray-900 dark:hover:text-white ms-5">
                  <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 17">
                    <path fill-rule="evenodd" d="M20 1.892a8.178 8.178 0 0 1-2.355.635 4.074 4.074 0 0 0 1.8-2.235 8.344 8.344 0 0 1-2.605.98A4.13 4.13 0 0 0 13.85 0a4.068 4.068 0 0 0-4.1 4.038 4 4 0 0 0 .105.919A11.705 11.705 0 0 1 1.4.734a4.006 4.006 0 0 0 1.268 5.392 4.165 4.165 0 0 1-1.859-.5v.05A4.057 4.057 0 0 0 4.1 9.635a4.19 4.19 0 0 1-1.856.07 4.108 4.108 0 0 0 3.831 2.807A8.36 8.36 0 0 1 0 14.184 11.732 11.732 0 0 0 6.291 16 11.502 11.502 0 0 0 17.964 4.5c0-.177 0-.35-.012-.523A8.143 8.143 0 0 0 20 1.892Z" clip-rule="evenodd"/>
                </svg>
                  <span class="sr-only">Twitter page</span>
              </a>
              <a href="#" class="text-gray-500 hover:text-gray-900 dark:hover:text-white ms-5">
                  <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M10 .333A9.911 9.911 0 0 0 6.866 19.65c.5.092.678-.215.678-.477 0-.237-.01-1.017-.014-1.845-2.757.6-3.338-1.169-3.338-1.169a2.627 2.627 0 0 0-1.1-1.451c-.9-.615.07-.6.07-.6a2.084 2.084 0 0 1 1.518 1.021 2.11 2.11 0 0 0 2.884.823c.044-.503.268-.973.63-1.325-2.2-.25-4.516-1.1-4.516-4.9A3.832 3.832 0 0 1 4.7 7.068a3.56 3.56 0 0 1 .095-2.623s.832-.266 2.726 1.016a9.409 9.409 0 0 1 4.962 0c1.89-1.282 2.717-1.016 2.717-1.016.366.83.402 1.768.1 2.623a3.827 3.827 0 0 1 1.02 2.659c0 3.807-2.319 4.644-4.525 4.889a2.366 2.366 0 0 1 .673 1.834c0 1.326-.012 2.394-.012 2.72 0 .263.18.572.681.475A9.911 9.911 0 0 0 10 .333Z" clip-rule="evenodd"/>
                  </svg>
                  <span class="sr-only">GitHub account</span>
              </a>
              <a href="#" class="text-gray-500 hover:text-gray-900 dark:hover:text-white ms-5">
                  <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M10 0a10 10 0 1 0 10 10A10.009 10.009 0 0 0 10 0Zm6.613 4.614a8.523 8.523 0 0 1 1.93 5.32 20.094 20.094 0 0 0-5.949-.274c-.059-.149-.122-.292-.184-.441a23.879 23.879 0 0 0-.566-1.239 11.41 11.41 0 0 0 4.769-3.366ZM8 1.707a8.821 8.821 0 0 1 2-.238 8.5 8.5 0 0 1 5.664 2.152 9.608 9.608 0 0 1-4.476 3.087A45.758 45.758 0 0 0 8 1.707ZM1.642 8.262a8.57 8.57 0 0 1 4.73-5.981A53.998 53.998 0 0 1 9.54 7.222a32.078 32.078 0 0 1-7.9 1.04h.002Zm2.01 7.46a8.51 8.51 0 0 1-2.2-5.707v-.262a31.64 31.64 0 0 0 8.777-1.219c.243.477.477.964.692 1.449-.114.032-.227.067-.336.1a13.569 13.569 0 0 0-6.942 5.636l.009.003ZM10 18.556a8.508 8.508 0 0 1-5.243-1.8 11.717 11.717 0 0 1 6.7-5.332.509.509 0 0 1 .055-.02 35.65 35.65 0 0 1 1.819 6.476 8.476 8.476 0 0 1-3.331.676Zm4.772-1.462A37.232 37.232 0 0 0 13.113 11a12.513 12.513 0 0 1 5.321.364 8.56 8.56 0 0 1-3.66 5.73h-.002Z" clip-rule="evenodd"/>
                </svg>
                  <span class="sr-only">Dribbble account</span>
              </a>
          </div>
      </div>
    </div>
</footer>


  <div class="bg-white">
    <div class="pt-6">
      <nav aria-label="Breadcrumb">
        <ol role="list" class="mx-auto flex max-w-2xl items-center space-x-2 px-4 sm:px-6 lg:max-w-7xl lg:px-8">
          <li v-for="breadcrumb in product.breadcrumbs" :key="breadcrumb.id">
            <div class="flex items-center">
              <a :href="breadcrumb.href" class="mr-2 text-sm font-medium text-gray-900">{{ breadcrumb.name }}</a>
              <svg width="16" height="20" viewBox="0 0 16 20" fill="currentColor" aria-hidden="true"
                class="h-5 w-4 text-gray-300">
                <path d="M5.697 4.34L8.98 16.532h1.327L7.025 4.341H5.697z" />
              </svg>
            </div>
          </li>
          <li class="text-sm">
            <a :href="product.href" aria-current="page" class="font-medium text-gray-500 hover:text-gray-600">{{
              product.name }}</a>
          </li>
        </ol>
      </nav>

      <!-- Image gallery -->
      <div class="mx-auto mt-6 max-w-2xl sm:px-6 lg:grid lg:max-w-7xl lg:grid-cols-3 lg:gap-x-8 lg:px-8">
        <img :src="product.images[0].src" :alt="product.images[0].alt"
          class="hidden aspect-[3/4] size-full rounded-lg object-cover lg:block" />
        <div class="hidden lg:grid lg:grid-cols-1 lg:gap-y-8">
          <img :src="product.images[1].src" :alt="product.images[1].alt"
            class="aspect-[3/2] size-full rounded-lg object-cover" />
          <img :src="product.images[2].src" :alt="product.images[2].alt"
            class="aspect-[3/2] size-full rounded-lg object-cover" />
        </div>
        <img :src="product.images[3].src" :alt="product.images[3].alt"
          class="aspect-[4/5] size-full object-cover sm:rounded-lg lg:aspect-[3/4]" />
      </div>

      <!-- Product info -->
      <div
        class="mx-auto max-w-2xl px-4 pb-16 pt-10 sm:px-6 lg:grid lg:max-w-7xl lg:grid-cols-3 lg:grid-rows-[auto,auto,1fr] lg:gap-x-8 lg:px-8 lg:pb-24 lg:pt-16">
        <div class="lg:col-span-2 lg:border-r lg:border-gray-200 lg:pr-8">
          <h1 class="text-2xl font-bold tracking-tight text-gray-900 sm:text-3xl">{{ product.name }}</h1>
        </div>

        <!-- Options -->
        <div class="mt-4 lg:row-span-3 lg:mt-0">
          <h2 class="sr-only">Product information</h2>
          <p class="text-3xl tracking-tight text-gray-900">{{ product.price }}</p>

          <!-- Reviews -->
          <div class="mt-6">
            <h3 class="sr-only">Reviews</h3>
            <div class="flex items-center">
              <div class="flex items-center">
                <StarIcon v-for="rating in [0, 1, 2, 3, 4]" :key="rating"
                  :class="[reviews.average > rating ? 'text-gray-900' : 'text-gray-200', 'size-5 shrink-0']"
                  aria-hidden="true" />
              </div>
              <p class="sr-only">{{ reviews.average }} out of 5 stars</p>
              <a :href="reviews.href" class="ml-3 text-sm font-medium text-indigo-600 hover:text-indigo-500">{{
                reviews.totalCount }} reviews</a>
            </div>
          </div>

          <form class="mt-10">
            <!-- Colors -->
            <div>
              <h3 class="text-sm font-medium text-gray-900">Color</h3>

              <fieldset aria-label="Choose a color" class="mt-4">
                <RadioGroup v-model="selectedColor" class="flex items-center space-x-3">
                  <RadioGroupOption as="template" v-for="color in product.colors" :key="color.name" :value="color"
                    :aria-label="color.name" v-slot="{ active, checked }">
                    <div
                      :class="[color.selectedClass, active && checked ? 'ring ring-offset-1' : '', !active && checked ? 'ring-2' : '', 'relative -m-0.5 flex cursor-pointer items-center justify-center rounded-full p-0.5 focus:outline-none']">
                      <span aria-hidden="true" :class="[color.class, 'size-8 rounded-full border border-black/10']" />
                    </div>
                  </RadioGroupOption>
                </RadioGroup>
              </fieldset>
            </div>

            <!-- Sizes -->
            <div class="mt-10">
              <div class="flex items-center justify-between">
                <h3 class="text-sm font-medium text-gray-900">Size</h3>
                <a href="#" class="text-sm font-medium text-indigo-600 hover:text-indigo-500">Size guide</a>
              </div>

              <fieldset aria-label="Choose a size" class="mt-4">
                <RadioGroup v-model="selectedSize" class="grid grid-cols-4 gap-4 sm:grid-cols-8 lg:grid-cols-4">
                  <RadioGroupOption as="template" v-for="size in product.sizes" :key="size.name" :value="size"
                    :disabled="!size.inStock" v-slot="{ active, checked }">
                    <div
                      :class="[size.inStock ? 'cursor-pointer bg-white text-gray-900 shadow-sm' : 'cursor-not-allowed bg-gray-50 text-gray-200', active ? 'ring-2 ring-indigo-500' : '', 'group relative flex items-center justify-center rounded-md border px-4 py-3 text-sm font-medium uppercase hover:bg-gray-50 focus:outline-none sm:flex-1 sm:py-6']">
                      <span>{{ size.name }}</span>
                      <span v-if="size.inStock"
                        :class="[active ? 'border' : 'border-2', checked ? 'border-indigo-500' : 'border-transparent', 'pointer-events-none absolute -inset-px rounded-md']"
                        aria-hidden="true" />
                      <span v-else aria-hidden="true"
                        class="pointer-events-none absolute -inset-px rounded-md border-2 border-gray-200">
                        <svg class="absolute inset-0 size-full stroke-2 text-gray-200" viewBox="0 0 100 100"
                          preserveAspectRatio="none" stroke="currentColor">
                          <line x1="0" y1="100" x2="100" y2="0" vector-effect="non-scaling-stroke" />
                        </svg>
                      </span>
                    </div>
                  </RadioGroupOption>
                </RadioGroup>
              </fieldset>
            </div>

            <button type="submit"
              class="mt-10 flex w-full items-center justify-center rounded-md border border-transparent bg-indigo-600 px-8 py-3 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
              Thêm vào giỏ hàng</button>
          </form>
        </div>

        <div class="py-10 lg:col-span-2 lg:col-start-1 lg:border-r lg:border-gray-200 lg:pb-16 lg:pr-8 lg:pt-6">
          <!-- Description and details -->
          <div>
            <h3 class="sr-only">Description</h3>

            <div class="space-y-6">
              <p class="text-base text-gray-900">{{ product.description }}</p>
            </div>
          </div>

          <div class="mt-10">
            <h3 class="text-sm font-medium text-gray-900">Thông tin sản phẩm</h3>

            <div class="mt-4">
              <ul role="list" class="list-disc space-y-2 pl-4 text-sm">
                <li v-for="information in product.informations" :key="information" class="text-gray-400">
                  <span class="text-gray-600">{{ information }}</span>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { StarIcon } from '@heroicons/vue/20/solid'
import { RadioGroup, RadioGroupOption } from '@headlessui/vue'
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
} from '@headlessui/vue'
import { Bars3Icon, MagnifyingGlassIcon, ShoppingBagIcon, XMarkIcon } from '@heroicons/vue/24/outline'
import { ChevronLeftIcon, ChevronRightIcon } from '@heroicons/vue/20/solid'
import { RouterLink } from 'vue-router'

const product = {
  name: 'ANGRY BIRDS Xe trớn tốc độ cao Angry Birds của chú chim nóng tính RED RED/AB23031',
  price: '$192',
  href: '#',
  breadcrumbs: [
    { id: 1, name: 'Trang chủ', href: '#' },
    { id: 2, name: 'Angry Bird', href: '#' },
  ],
  images: [
    {
      src: "./src/assets/image/AngryBird-1.webp",
      alt: 'Two each of gray, white, and black shirts laying flat.',
    },
    {
      src: "./src/assets/image/AngryBird-2.webp",
      alt: 'Model wearing plain black basic tee.',
    },
    {
      src: "./src/assets/image/AngryBird-3.webp",
      alt: 'Model wearing plain gray basic tee.',
    },
    {
      src: "./src/assets/image/AngryBird-4.webp",
      alt: 'Model wearing plain white basic tee.',
    },
  ],
  colors: [
    { name: 'White', class: 'bg-white', selectedClass: 'ring-gray-400' },
    { name: 'Gray', class: 'bg-gray-200', selectedClass: 'ring-gray-400' },
    { name: 'Black', class: 'bg-gray-900', selectedClass: 'ring-gray-900' },
  ],
  sizes: [
    { name: 'XXS', inStock: false },
    { name: 'XS', inStock: true },
    { name: 'S', inStock: true },
    { name: 'M', inStock: true },
    { name: 'L', inStock: true },
    { name: 'XL', inStock: true },
    { name: '2XL', inStock: true },
    { name: '3XL', inStock: true },
  ],
  description:
  'Angry Birds Crashers là những chiếc xe trớn siêu nhanh với thiết kế cool ngầu. '+
    'Mô phỏng các nhân vật trong phim hoạt hình Angry Birds. Có 6 nhân vật: Red, Leonard, Hog, Bomb, Terence, Chuck. '+
    'Bố mẹ hãy sưu tập đủ các nhân vật để giúp các bé có thể viết lên những câu truyện như trong phim nhé!',
  informations: [
    'Chủ đề: ANGRY BLISTER',
    'Xuất xứ: Trung Quốc',
    'Mã VT: RED/AB23031',
    'Thương hiệu: Angry Bird',
  ],
}
const reviews = { href: '#', average: 4, totalCount: 117 }

const selectedColor = ref(product.colors[0])
const selectedSize = ref(product.sizes[2])


const navigation = {
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

const products = [
  {
    id: 1,
    name: 'Minions',
    href: '#',
    price: '569.000đ',
    imageSrc: 'https://www.mykingdom.com.vn/cdn/shop/files/mo-hinh-nhan-vat-diep-vien-minions-desicapable-me-4-minions-gj5t.png?v=1732850490&width=713',
    imageAlt: 'Tall slender porcelain bottle with natural clay textured body and cork stopper.',
  },
  {
    id: 2,
    name: 'PopBean Lucky Cat Series',
    href: '#',
    price: '240.000đ',
    imageSrc: 'https://www.mykingdom.com.vn/cdn/shop/files/mo-hinh-pop-bean-lucky-cat-series-pop-mart-6931571007287_0.jpg?v=1732681931&width=713',
    imageAlt: 'Olive drab green insulated bottle with flared screw lid and flat top.',
  },
  {
    id: 3,
    name: 'Mô hình MollyCarb-LoverSeries',
    href: '#',
    price: '280.000đ',
    imageSrc: 'https://www.mykingdom.com.vn/cdn/shop/files/mo-hinh-molly-carb-lover-series-figures-pop-mart-6931571006662_0.jpg?v=1732680845&width=713',
    imageAlt: 'Person using a pen to cross a task off a productivity paper card.',
  },
  {
    id: 4,
    name: 'Đồ chơi lắp ráp Diy Xe Van',
    href: '#',
    price: '419.000đ',
    imageSrc: 'https://www.mykingdom.com.vn/cdn/shop/files/do-choi-lap-rap-diy-xe-van-cap-cuu-co-den-va-am-thanh-vecto-vt8088c_1.jpg?v=1732246150&width=713',
    imageAlt: 'Hand holding black machined steel mechanical pencil with brass tip and top.',
  },
  // More products...
]



const items = [
  { id: 1, title: 'Back End Developer', department: 'Engineering', type: 'Full-time', location: 'Remote' },
  { id: 2, title: 'Front End Developer', department: 'Engineering', type: 'Full-time', location: 'Remote' },
  { id: 3, title: 'User Interface Designer', department: 'Design', type: 'Full-time', location: 'Remote' },
]
const open = ref(false)




</script>