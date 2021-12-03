# frozen_string_literal: true

class CashesController < ApplicationController
  def update
    render :edit, status: :unprocessable_entity unless resource.update(resource_params)
  end

  private

  def cashes
    Cash.order(:name)
  end

  def collection
    @collection ||= CashSearcher.search(cashes, params)
  end

  def resource_params
    params.require(:cash).permit(:formula, :name, :currency)
  end

  def resource
    @resource ||= ::Cashes::GetResourceService.call(params)
  end
end
