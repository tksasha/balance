# frozen_string_literal: true

class ItemDecorator < Draper::Decorator
  delegate_all

  def date
    helpers.localize(model.date) if model.date.is_a?(Date)
  end

  def description
    return unless model.description.respond_to?(:gsub)

    #
    # converts
    # "[First Tag] [Second Tag] some description"
    # to
    # "<div class="tag">First Tag</div> <div class="tag">Second Tag</div> some description"
    #
    model
      .description
      .gsub(/\[([[[:alnum:]][[:blank:]]']+)\]/, '<div class="tag">\1</div>')
  end
end
