module Permitter
  def acp params
    ActionController::Parameters.new params
  end

  def permit! params
    acp(params).permit!
  end
end
