# frozen_string_literal: true

class CashesController < InheritedResources::Base
  respond_to :js

  def update
    update! do |_, failure|
      failure.js { render :edit }
    end
  end

  private

  def resource_params
    params.require(:cash).permit(:name, :formula)
  end

  # TODO: spec me?
  def collection
    @collection ||= CashSearcher.search(Cash.order(:name), params)
  end
end
