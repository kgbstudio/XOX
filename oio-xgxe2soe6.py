class BotSearch:
    def __init__(self, data):
        self.data = data

    def search(self, query):
        results = []
        for item in self.data:
            if query.lower() in item.lower():
                results.append(item)
        return results

data = ["apple", "banana", "cherry", "date", "elderberry"]
bot = BotSearch(data)
print(bot.search("a")) 
print(bot.search("b")) 
print(bot.search("c")) 
print(bot.search("d")) 
print(bot.search("e")) 
print(bot.search("f"))