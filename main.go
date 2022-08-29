
func add_phantomwallet(geturls string) {

	fmt.Println("starting...")
	extPath, _ := filepath.Abs("fixtures/chrome-extension")

	u := launcher.New().
		// Must use abs path for an extension
		Set("load-extension", extPath).
		Headless(false).Devtools(false).MustLaunch()

	page := rod.New().ControlURL(u).MustConnect().MustPage("chrome-extension://bfnaelmomeimhlpmgjnjophhpkkoljpa/onboarding.html").MustWaitLoad() //kan ikke ændres
	closefirst, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	closefirst.HasCTRL(true) //ctrl

	closefirst.SetKeys(keybd_event.VK_W) // w knap
	closefirst.Press()
	closefirst.Release() 


	page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > div > div.sc-iCfMLu.sc-dvQaRk.hJnMKN.dsAYhQ > button.sc-eCImPb.fimA-Dk").MustClick()
	page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-jeraig.dThhlm.bgUxUa > input").MustInput("antibot1234").MustType(input.Enter)
	page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-jeraig.dThhlm.bgUxUa > div > div > input").MustInput("antibot1234").MustType(input.Enter)
	page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-bdvvtL.sc-hBUSln.bdpdgN.ciZnfo > span > input[type=checkbox]").MustClick()
	page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > button").MustClick() //continue on password

	page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx").MustClick()
	//getkeyphrase := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd ").MustText()
	getkeyphrase1 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(1) > input").MustText()

	getkeyphrase2 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(2) > input").MustText()

	getkeyphrase3 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(3) > input").MustText()

	getkeyphrase4 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(4) > input").MustText()
	getkeyphrase5 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(5) > input").MustText()
	getkeyphrase6 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(6) > input").MustText()
	getkeyphrase7 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(7) > input").MustText()
	getkeyphrase8 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(8) > input").MustText()
	getkeyphrase9 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(9) > input").MustText()
	getkeyphrase10 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(10) > input").MustText()
	getkeyphrase11 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(11) > input").MustText()
	getkeyphrase12 := page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-iCfMLu.sc-hUpaCq.chZIgj.daGzFx > div > div.sc-cNKqjZ.jPlcxd > div:nth-child(12) > input").MustText()

	//fmt.Println(getkeyphrase1, getkeyphrase2, getkeyphrase3, getkeyphrase4, getkeyphrase5, getkeyphrase6, getkeyphrase7, getkeyphrase8, getkeyphrase9, getkeyphrase10, getkeyphrase11, getkeyphrase12)
	fmt.Println("trying to create json")

	fmt.Println(getkeyphrase1, getkeyphrase2, getkeyphrase3, getkeyphrase4, getkeyphrase5, getkeyphrase6, getkeyphrase7, getkeyphrase8, getkeyphrase9, getkeyphrase10, getkeyphrase11, getkeyphrase12)

//make your own json here.

	bytes, _ := json.MarshalIndent(ta, "", "\t")
	fmt.Println(string(bytes))
	fmt.Println("----------------------------------")
	enc := json.NewEncoder(os.Stdout)
	fmt.Println("json encoded successfully")
	if err := enc.Encode(ta); err != nil {
		log.Fatal(err)
	}

	

	_ = ioutil.WriteFile("solanasniper.json", file, 0644)
	fmt.Println("writed json successfully")
	utils.Sleep(1000)
	//see keyphrase
	page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > div > div > div.sc-bdvvtL.sc-hBUSln.bdpdgN.ciZnfo > span > input[type=checkbox]").MustClick() //save keyphrase
	page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > button").MustClick()
	page.MustElement("#root > main > div.sc-gIDmLj.hNrXGi > form > button").MustClick()
	fmt.Println("creating a new wallet successfully")

	//error

	page.MustEval(`() => console.log("initialized")`)
	fmt.Println(page.MustInfo().Title)
	fmt.Println("initialized...")
	fmt.Println("waiting for a desired task please wait...")

	//her
	//urlgetter()
	//list[0]
	page1 := rod.New().ControlURL(u).MustConnect().MustPage(geturls)

	fmt.Println("success")
	fmt.Println("worker successfully waiting to start task...")

	/*






	 */
	page1 = rod.New().ControlURL(u).MustConnect().MustPage("https://magiceden.io/")
	page1.MustElement("#root > div.main.page.tw-bg-\\[\\#120d18\\] > header > nav > div.tw-flex.tw-items-center.tw-space-x-2 > div.tw-hidden.md\\:tw-inline-flex.tw-items-center.tw-justify-end.tw-space-x-4 > div > button:nth-child(2)").MustClick()
	page1.MustElement("#root > div.modal.fade.show > div > div > div.modal-body.tw-pt-0 > ul > li.tw-border-pink-200.tw-p-2.hover\\:tw-bg-pink-200.hover\\:tw-bg-opacity-10.hover\\:tw-border-pink-200.tw-border-2.tw-transition-all.tw-rounded-lg > button").MustClick()
	page1.MustElement("#root > div.main.page.tw-bg-\\[\\#120d18\\] > header > nav > div.tw-flex.tw-items-center.tw-space-x-2 > div.tw-hidden.md\\:tw-inline-flex.tw-items-center.tw-justify-end.tw-space-x-4 > div > div > div > div.cursor-pointer.position-relative > button > span > span").MustClick()
	page1.MustElement("#root > div.main.page.tw-bg-\\[\\#120d18\\] > header > nav > div.tw-flex.tw-items-center.tw-space-x-2 > div.tw-hidden.md\\:tw-inline-flex.tw-items-center.tw-justify-end.tw-space-x-4 > div > div > div > div.dropdown.tw-text-secondary.no-border.visible > div > div.tw-flex.tw-m-4.tw-space-x-2 > div > a").MustClick()
	//fmt.Println(page1.MustElement("#content > section > div.tw-flex.tw-flex-auto.tw-items-center.md\\:tw-items-start.tw-flex-col.md\\:tw-flex-row.tw-px-4.sm\\:tw-pl-8.md\\:tw-pt-16.lg\\:tw-pl-10.\\32 xl\\:tw-pl-12.tw-pb-8.md\\:tw-pb-24 > div.tw-grid.tw-grid-cols-1.\\32 xl\\:tw-grid-cols-2.tw-mt-5.md\\:tw-mt-0.md\\:tw-ml-6.lg\\:tw-ml-14.\\32 xl\\:tw-ml-24 > div.tw-mt-1.md\\:tw-mt-4.tw-space-y-7.tw-flex.tw-flex-col.sm\\:tw-items-center.md\\:tw-items-start.tw-row-start-2 > div > div.tw-flex.tw-mb-8.tw-gap-3 > button.tw-inline-flex.tw-justify-center.tw-items-center.tw-rounded-md.tw-text-white-1.CopiableButton_btn__oahaK.tw-flex.tw-items-center.tw-space-x-2.tw-border.tw-border-gray-400.tw-px-2\\.5.tw-py-1.\\!tw-rounded-\\[200px\\].tw-text-white-2.tw-text-sm.hover\\:tw-bg-gray-200").MustEval(`() => document.querySelector('button[title]');`).String())
	text := page1.MustElement("#content > section > div.tw-flex.tw-flex-auto.tw-items-center.md\\:tw-items-start.tw-flex-col.md\\:tw-flex-row.tw-px-4.sm\\:tw-pl-8.md\\:tw-pt-16.lg\\:tw-pl-10.\\32 xl\\:tw-pl-12.tw-pb-8.md\\:tw-pb-24 > div.tw-grid.tw-grid-cols-1.\\32 xl\\:tw-grid-cols-2.tw-mt-5.md\\:tw-mt-0.md\\:tw-ml-6.lg\\:tw-ml-14.\\32 xl\\:tw-ml-24 > div.tw-mt-1.md\\:tw-mt-4.tw-space-y-7.tw-flex.tw-flex-col.sm\\:tw-items-center.md\\:tw-items-start.tw-row-start-2 > div > div.tw-flex.tw-mb-8.tw-gap-3 > button.tw-inline-flex.tw-justify-center.tw-items-center.tw-rounded-md.tw-text-white-1.CopiableButton_btn__oahaK.tw-flex.tw-items-center.tw-space-x-2.tw-border.tw-border-gray-400.tw-px-2\\.5.tw-py-1.\\!tw-rounded-\\[200px\\].tw-text-white-2.tw-text-sm.hover\\:tw-bg-gray-200").MustEval(`() => console.log(document.querySelectorAll('button[title]'));`).String()

	val := page1.MustEval(`() =>  console.log(document.querySelectorAll('button[title]').forEach(btn => {
  let caption = document.createElement('div')
  caption.textContent = btn.title
  btn.appendChild(caption)
}));`).Get("title").Str()
	fmt.Println(val, "fixed content of ME to get wallet adress") // output: jack
	//burner := page1.MustElement("#content > section > div.tw-flex.tw-flex-auto.tw-items-center.md\\:tw-items-start.tw-flex-col.md\\:tw-flex-row.tw-px-4.sm\\:tw-pl-8.md\\:tw-pt-16.lg\\:tw-pl-10.\\32 xl\\:tw-pl-12.tw-pb-8.md\\:tw-pb-24 > div.tw-grid.tw-grid-cols-1.\\32 xl\\:tw-grid-cols-2.tw-mt-5.md\\:tw-mt-0.md\\:tw-ml-6.lg\\:tw-ml-14.\\32 xl\\:tw-ml-24 > div.tw-mt-1.md\\:tw-mt-4.tw-space-y-7.tw-flex.tw-flex-col.sm\\:tw-items-center.md\\:tw-items-start.tw-row-start-2 > div > div.tw-flex.tw-mb-8.tw-gap-3 > button.tw-inline-flex.tw-justify-center.tw-items-center.tw-rounded-md.tw-text-white-1.CopiableButton_btn__oahaK.tw-flex.tw-items-center.tw-space-x-2.tw-border.tw-border-gray-400.tw-px-2\\.5.tw-py-1.\\!tw-rounded-\\[200px\\].tw-text-white-2.tw-text-sm.hover\\:tw-bg-gray-200").MustClick().MustText()
	fmt.Println("--------------------------------")
	//fmt.Println(burner)
	test := page1.MustElement("#content > section > div.tw-flex.tw-flex-auto.tw-items-center.md\\:tw-items-start.tw-flex-col.md\\:tw-flex-row.tw-px-4.sm\\:tw-pl-8.md\\:tw-pt-16.lg\\:tw-pl-10.\\32 xl\\:tw-pl-12.tw-pb-8.md\\:tw-pb-24 > div.tw-grid.tw-grid-cols-1.\\32 xl\\:tw-grid-cols-2.tw-mt-5.md\\:tw-mt-0.md\\:tw-ml-6.lg\\:tw-ml-14.\\32 xl\\:tw-ml-24 > div.tw-mt-1.md\\:tw-mt-4.tw-space-y-7.tw-flex.tw-flex-col.sm\\:tw-items-center.md\\:tw-items-start.tw-row-start-2 > div > div.tw-flex.tw-mb-8.tw-gap-3 > button.tw-inline-flex.tw-justify-center.tw-items-center.tw-rounded-md.tw-text-white-1.CopiableButton_btn__oahaK.tw-flex.tw-items-center.tw-space-x-2.tw-border.tw-border-gray-400.tw-px-2\\.5.tw-py-1.\\!tw-rounded-\\[200px\\].tw-text-white-2.tw-text-sm.hover\\:tw-bg-gray-200 > div:nth-child(3)").MustText()
	fmt.Println(test, "\n your burner wallet adress")
	fmt.Println(text)

	fmt.Println("--------------------------------")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("what price you seek?", "\n safe the key now so you dont loose your solana you send. it can crash!\n", "write the desired amount: ")
	price, _ := reader.ReadString('\n')

	if price == "" {

		utils.Sleep(5)

	} else {
		goto begin
	}
begin:
	i := 1
	max := 2000000
	for i < max {

		utils.Sleep(0.5)
		fmt.Println(i)
		i += 1
		browser := rod.New().Timeout(time.Minute).MustConnect()
		defer browser.MustClose()

		// You can also use stealth.JS directly without rod
		fmt.Printf("js: %x\n\n", md5.Sum([]byte(stealth.JS)))

		page := stealth.MustPage(browser)

		page.MustNavigate("https://api.coralcube.io/v1/getItems?offset=0&page_size=24&ranking=price_asc&symbol=elenawhaleclub").MustWaitLoad()

		rand.Seed(time.Now().UnixNano())

		fmt.Println(randSeq(10), "unica id")

		text := page.MustElement("body > pre").MustText()
		var jsonoutput = text

		fmt.Println("parsed")
		fmt.Println("starting to check if its valid.")

		result := gjson.Get(jsonoutput, `items.#.name.@reverse|0`)
		result.ForEach(func(key, value gjson.Result) bool {
			println(value.String())
			return true // keep iterating
		})

		floorprice := gjson.Get(jsonoutput, `items.#.price.@reverse|1`)
		under_floorprice := gjson.Get(jsonoutput, `items.#.price.@reverse|0`)
		geturls := gjson.Get(jsonoutput, `items.#.buy_link.@reverse|0`)

		// The amount to send (in lamports);
		// 1 sol = 1000000000 lamports

		println(cmp.Equal(under_floorprice, floorprice))

		if cmp.Equal(under_floorprice, floorprice) { //checks if its equal or not
			fmt.Println(cmp.Equal(under_floorprice, floorprice))

			fmt.Println("floor price hasnt moved.")
		} else {
			fmt.Println(cmp.Equal(under_floorprice, floorprice))
			fmt.Println(under_floorprice, floorprice)
			page1 = rod.New().ControlURL(u).MustConnect().MustPage(geturls.String())
			list = append(list, geturls.String())
			break

		}

		if i == max {
			fmt.Println("Finished")
			break
		}
	}
