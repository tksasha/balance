# frozen_string_literal: true

class MoneyDecorator < Draper::Decorator
  def to_s
    helpers.number_with_delimiter format('%.2f', object)
  end
end
