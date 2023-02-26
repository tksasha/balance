# frozen_string_literal: true

RSpec.describe ConsolidationsController do
  it { is_expected.to be_a(BaseController) }

  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { is_expected.to eq :collection }
    end

    context do
      before { allow(subject).to receive_message_chain(:dashboard, :consolidations).and_return(:collection) }

      its(:collection) { is_expected.to eq :collection }
    end
  end
end
