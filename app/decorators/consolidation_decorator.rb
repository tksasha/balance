# frozen_string_literal: true

class ConsolidationDecorator < Draper::Decorator
  delegate_all

  delegate :year, :month, to: :date

  private

  def date
    context[:date]
  end
end
