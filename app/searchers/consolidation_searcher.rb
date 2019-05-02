class ConsolidationSearcher
  def initialize relation, params
    @relation = relation

    @date = params[:date]
  end

  def search
    @relation.where(date: date_range).select('SUM(sum) AS sum, category_id').group(:category_id).tap do |items|
      ConsolidationExpensesSum.sum = items.select { |item| item.income? == false }.map(&:sum).sum
    end
  end

  private
  def date_range
    DateRange.month @date
  end

  class << self
    def search *args
      new(*args).search
    end
  end
end
