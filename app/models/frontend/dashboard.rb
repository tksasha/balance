# frozen_string_literal: true

module Frontend
  class Dashboard
    def initialize(params = {})
      @currency = Currency.parse(params[:currency])

      @date = DateRange.parse(params)
    end

    def items
      scope
        .where(date:)
        .order(date: :desc)
    end

    def income
      scope
        .income
        .sum(:sum)
    end

    def expense
      scope
        .expense
        .sum(:sum)
    end

    def at_end
      (income - expense).round(2)
    end

    private

    attr_reader :currency, :date

    def scope
      Item.where(currency:)
    end
  end
end
