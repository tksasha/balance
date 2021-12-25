# frozen_string_literal: true

module MoneyHelper
  def money(sum)
    MoneyDecorator.new(sum).to_s
  end
end
