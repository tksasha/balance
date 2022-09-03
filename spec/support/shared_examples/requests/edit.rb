# frozen_string_literal: true

RSpec.shared_examples 'edit.js' do
  it { is_expected.to render_template :edit }

  it { expect(response.content_type).to eq 'text/javascript; charset=utf-8' }
end
