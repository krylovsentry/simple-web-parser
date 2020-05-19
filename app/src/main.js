document.addEventListener("DOMContentLoaded", () => {
    new Vue({
        el: 'body',
        data: {
            news: [],
            newFeed: {},
            search: "",
            notFound: "Nothing found"
        },
        methods: {
            addFeeder: function() {
                if (!Object.values(this.newFeed).every(value => value)) {
                    alert("Fill all values");
                } else {
                    this.$http.post('/feeder', this.newFeed).success(function(response) {
                        this.newFeed = {}
                    }).error(function(error) {
                    });
                }
            },
            searchNews: function () {
                if (this.search && (this.search.length < 3)) {
                    alert('Please enter more than 3 symbols');
                } else {
                    this.$http.get(`/news?search=${this.search}`).success(function(response) {
                        this.news = response.items || [];
                        this.search = "";
                        if (!this.news.length) {
                            this.notFound = "Nothing found";
                        } else {
                            this.notFound = "";
                        }
                    });
                }
            },
            fill: function () {
                if (Math.random() < 0.5) {
                    this.newFeed = {
                        url: "https://vc.ru/",
                        oneParent: "feed__item",
                        content: "content",
                        title: "content-header"
                    }
                } else {
                    this.newFeed = {
                        url: "https://news.ycombinator.com/newest",
                        oneParent: "athing",
                        content: "content",
                        title: "title"
                    }
                }
            }
        }
    })
});

