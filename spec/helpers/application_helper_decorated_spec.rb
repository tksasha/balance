# frozen_string_literal: true

RSpec.describe ApplicationHelper do
  subject { helper }

  describe '#decorated' do
    subject { helper.decorated }

    before { allow(helper).to receive_message_chain(:resource, :decorate).and_return(:decorated) }

    it { is_expected.to eq :decorated }
  end
end
