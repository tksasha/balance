# frozen_string_literal: true

module Frontend
  class Dashboard
    def initialize(params)
      @currency = params[:currency]

      @date = Date.new(params[:year], params[:month], 1).all_month
    end

    def items
      Item
        .where(currency:, date:)
        .order(date: :desc)
    end

    private

    attr_reader :currency, :date
  end
end
