# frozen_string_literal: true

class ItemDecorator < Draper::Decorator
  delegate_all

  def date
    model.date.strftime '%d.%m.%Y' if model.date.respond_to? :strftime
  end
end
